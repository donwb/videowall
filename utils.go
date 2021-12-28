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

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println("--------------------")
		fmt.Println("PANIC: ", msg)
		panic(err)
	}
}
