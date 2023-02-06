package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type ThumbnailResponse struct {
	Data []struct {
		TargetID int64  `json:"targetId"`
		State    string `json:"state"`
		ImageURL string `json:"imageUrl"`
	} `json:"data"`
}

// Cache thumbnails to avoid spamming the API
var cache = make(map[int64][]byte)
var cacheExpiration = make(map[int64]time.Time)

// Download a user's thumbnail and expire the cache after 5 minutes
func downloadThumbnail(userID int64) (io.Reader, error) {
	now := time.Now()
	if thumbnailData, ok := cache[userID]; ok {
		if now.Before(cacheExpiration[userID]) {
			return bytes.NewReader(thumbnailData), nil
		}
	}

	// https://thumbnails.roblox.com/docs/index.html
	resp, err := http.Get(fmt.Sprintf("https://thumbnails.roblox.com/v1/users/avatar-headshot?userIds=%d&size=352x352&format=Png&isCircular=false", userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var thumbnailResponse ThumbnailResponse
	err = json.Unmarshal(body, &thumbnailResponse)
	if err != nil {
		return nil, err
	}

	if len(thumbnailResponse.Data) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	thumbnail := thumbnailResponse.Data[0]
	if thumbnail.State != "Completed" {
		return nil, fmt.Errorf("thumbnail not ready")
	}

	thumbnailURL, err := url.Parse(thumbnail.ImageURL)
	if err != nil {
		return nil, err
	}

	thumbnailURL.Scheme = "https"
	thumbnailURL.Host = "tr.rbxcdn.com"

	log.Printf("Downloading thumbnail from %s\n", thumbnailURL.String())
	thumbnailResp, err := http.Get(thumbnailURL.String())
	if err != nil {
		return nil, err
	}

	thumbnailData, err := io.ReadAll(thumbnailResp.Body)
	if err != nil {
		return nil, err
	}
	defer thumbnailResp.Body.Close()

	cache[userID] = thumbnailData
	cacheExpiration[userID] = now.Add(time.Minute * 5)

	return bytes.NewReader(thumbnailData), nil
}
