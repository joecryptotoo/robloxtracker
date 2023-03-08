package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gregdel/pushover"
)

func notifyPresenceChange(user User) {
	app := pushover.New(os.Getenv("PUSHOVER_APP_TOKEN"))
	recipient := pushover.NewRecipient(os.Getenv("PUSHOVER_USER_KEY"))

	minutesSinceLastState := int(time.Now().UTC().Sub(user.LastPresenceChange).Minutes())

	minutesSinceLastOnline := int(time.Now().UTC().Sub(user.Presence.LastOnline).Minutes())

	lastPresenceType := presenceTypeToString(user.LastPresenceType)
	presenceType := presenceTypeToString(user.Presence.UserPresenceType)
	message := pushover.NewMessage(fmt.Sprintf("User %s is now %s, was %s for %d minutes, last online %d minutes ago.",
		user.Name, presenceType, lastPresenceType, minutesSinceLastState, minutesSinceLastOnline))
	message.Title = "Roblox Presence Change"
	message.URL = fmt.Sprintf("https://www.roblox.com/users/%d/profile", user.ID)
	message.URLTitle = "View Profile"
	message.Sound = "magic"

	thumbnailData, err := downloadThumbnail(user.ID)
	if err != nil {
		log.Println(err)
	} else {
		message.AddAttachment(thumbnailData)
	}

	// Update metrics
	user.Metrics.UserPresenceType.Set(float64(user.LastPresenceType))

	switch user.LastPresenceType {
	case 0:
		user.Metrics.OfflineTime.Set(float64(minutesSinceLastState))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 1:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(minutesSinceLastState))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 2:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(minutesSinceLastState))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(0))
	case 3:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(minutesSinceLastState))
		user.Metrics.UnknownTime.Set(float64(0))
	default:
		user.Metrics.OfflineTime.Set(float64(0))
		user.Metrics.OnlineTime.Set(float64(0))
		user.Metrics.InGameTime.Set(float64(0))
		user.Metrics.InStudioTime.Set(float64(0))
		user.Metrics.UnknownTime.Set(float64(minutesSinceLastState))
	}

	log.Printf("Presence state changed to %s. Notifying. \"%s\"\n", presenceTypeToString(user.Presence.UserPresenceType), message.Message)

	_, err = app.SendMessage(message, recipient)
	if err != nil {
		log.Println(err)
	}
}
