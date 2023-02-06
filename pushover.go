package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gregdel/pushover"
)

func notifyPresenceChange(user User) {
	app := pushover.New(os.Getenv("PUSHOVER_APP_TOKEN"))
	recipient := pushover.NewRecipient(os.Getenv("PUSHOVER_USER_KEY"))

	presenceType := presenceTypeToString(user.Presence.UserPresenceType)
	message := pushover.NewMessage(fmt.Sprintf("User %s is now %s", user.Name, presenceType))
	message.Title = "Roblox Presence Change"
	message.URL = fmt.Sprintf("https://www.roblox.com/users/%d/profile", user.ID)
	message.URLTitle = "View Profile"
	message.Sound = "magic"

	thumbnailData, err := downloadThumbnail(user.ID)
	if err != nil {
		log.Println(err)
		return
	}

	message.AddAttachment(thumbnailData)

	log.Printf("Presence state changed to %s. Notifying. \"%s\"\n", presenceTypeToString(user.Presence.UserPresenceType), message.Message)
	_, err = app.SendMessage(message, recipient)
	if err != nil {
		log.Println(err)
	}
}
