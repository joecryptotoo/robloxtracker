package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	// Create a non-global registry.
	reg := prometheus.NewRegistry()

	// Check for required environment variables
	userID, err := strconv.ParseInt(os.Getenv("ROBLOX_USER_ID"), 10, 64)
	if err != nil {
		log.Println(err)
		return
	}

	// Get the username string
	user, err := getUsernameFromID(userID)
	if err != nil {
		log.Println(err)
		return
	}

	user.LastPresenceChange = time.Now().UTC()

	user.Metrics = *RobloxMetrics(reg)

	user.Metrics.OfflineTime.Set(float64(0))
	user.Metrics.OnlineTime.Set(float64(0))
	user.Metrics.InGameTime.Set(float64(0))
	user.Metrics.InStudioTime.Set(float64(0))
	user.Metrics.UnknownTime.Set(float64(0))

	// Start presence checker
	presenceState := 0
	user.LastPresenceType = presenceState
	t := time.NewTicker(time.Second * 5)

	// Expose metrics and custom registry via an HTTP server
	// using the HandleFor function. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Check presence every 5 seconds
	for range t.C {
		// Check presence
		user.Presence, err = checkPresence(user.ID)
		if err != nil {
			log.Println(err)
			return
		}

		// Check if presence has changed and notify
		if presenceState != user.Presence.UserPresenceType {
			// Update last online time
			user.LastPresenceType = presenceState
			minutesSinceLastOnline := int(time.Now().UTC().Sub(user.Presence.LastOnline).Minutes())

			// Log presence change
			log.Printf("User %s is %s, last online %d minutes ago\n", user.Name, presenceTypeToString(user.Presence.UserPresenceType), minutesSinceLastOnline)

			log.Printf("Presence: %#v\n", user.Presence)

			// Notify if user is online
			notifyPresenceChange(user)
			user.LastPresenceChange = time.Now().UTC()
		}

		updateMetrics(user)

		// Update presence state
		presenceState = user.Presence.UserPresenceType
	}
}
