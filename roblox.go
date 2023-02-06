package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UserPresence struct {
	UserPresenceType int
	LastOnline       time.Time
	PlaceID          int64
	RootPlaceID      int64
	GameID           string
	UniverseID       int64
	UserID           int64
}

type UserPresenceResponse struct {
	UserPresences []UserPresence
}

type User struct {
	Description            string       `json:"description"`
	Created                string       `json:"created"`
	IsBanned               bool         `json:"isBanned"`
	ExternalAppDisplayName string       `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool         `json:"hasVerifiedBadge"`
	ID                     int64        `json:"id"`
	Name                   string       `json:"name"`
	DisplayName            string       `json:"displayName"`
	Presence               UserPresence `json:"userPresence"`
}

func GetUsernameFromID(id int64) (User, error) {
	resp, err := http.Get(fmt.Sprintf("https://users.roblox.com/v1/users/%d", id))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func checkPresence(userID int64) (UserPresence, error) {
	requestBody := struct {
		UserIDs []int64 `json:"userIds"`
	}{
		UserIDs: []int64{userID},
	}

	reqBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error encoding request body:", err)
		return UserPresence{}, err
	}

	resp, err := http.Post("https://presence.roblox.com/v1/presence/users", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		fmt.Println("Error making request:", err)
		return UserPresence{}, err
	}
	defer resp.Body.Close()

	var presenceResponse UserPresenceResponse
	err = json.NewDecoder(resp.Body).Decode(&presenceResponse)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return UserPresence{}, err
	}

	if len(presenceResponse.UserPresences) == 0 {
		fmt.Println("User not found")
		return UserPresence{}, nil
	}

	presence := presenceResponse.UserPresences[0]
	return presence, nil
}

func presenceTypeToString(presenceType int) string {
	switch presenceType {
	case 0:
		return "Offline"
	case 1:
		return "Online"
	case 2:
		return "InGame"
	case 3:
		return "InStudio"
	default:
		return "Unknown"
	}
}
