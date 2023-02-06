package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	presenceState := 0
	t := time.NewTicker(time.Second * 5)

	userID, err := strconv.ParseInt(os.Getenv("ROBLOX_USER_ID"), 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := getUsernameFromID(userID)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		select {
		case <-t.C:

			user.Presence, err = checkPresence(user.ID)
			if err != nil {
				log.Println(err)
				return
			}

			// Check if presence has changed and notify
			if presenceState != user.Presence.UserPresenceType {
				now := time.Now().UTC()
				minutesSinceLastOnline := int(now.Sub(user.Presence.LastOnline).Minutes())
				log.Printf("User %s is %s, last online %d minutes ago\n", user.Name, presenceTypeToString(user.Presence.UserPresenceType), minutesSinceLastOnline)

				notifyPresenceChange(user)
			}

			// Update presence state
			presenceState = user.Presence.UserPresenceType
		}
	}
}
