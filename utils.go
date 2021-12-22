package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

func calcLastTrackTime(lastPlayTime string) int {
	// get current unix time
	now := time.Now().UTC().Unix()
	fmt.Println("Current UNIX Time ", now)

	lastDateInt, err := strconv.Atoi(lastPlayTime)
	if err != nil {
		return 0
	}

	// div by 60 to get minutes
	dif := now - int64(lastDateInt)
	retVal := dif / 60

	return int(retVal)
}

func isNowPlaying(trackInfo lastfm.UserGetRecentTracks) bool {
	// get most recent track
	track := trackInfo.Tracks[0]
	isPlaying := track.NowPlaying == "true"

	if !isPlaying {
		trackUTS := track.Date.Uts
		minutesSinceLastTrack := calcLastTrackTime(trackUTS)
		return minutesSinceLastTrack < 10
	} else {
		return isPlaying
	}
}

func getAlbumArtAndPlayCount(artist string, track string) (string, string) {
	// refactor this out!
	apikey := apiKey
	apisecret := apiPassword

	api := lastfm.New(apikey, apisecret)

	// Make this checking better
	trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": track})

	playcount := trackInfo.PlayCount

	images := trackInfo.Album.Images
	if len(images) == 0 {
		newTrackName := removeRemastered(track)
		trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": newTrackName})
		fmt.Println("NEW TRACK NAME: ", newTrackName)
		if len(images) == 0 {
			return "placeholder", playcount
		} else {
			images = trackInfo.Album.Images
			return images[2].Url, playcount
		}

	}

	fmt.Println("Images: ", images)

	largeImage := images[2].Url
	return largeImage, playcount

	//return "https://lastfm.freetls.fastly.net/i/u/174s/e5078801aed03ec9bc933a736349f143.png"
}

func removeRemastered(trackName string) string {
	// try to remove the (remastered) from the end of a song if it gets in the way of album art
	return "Champagne Supernova"
}
