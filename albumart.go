package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shkh/lastfm-go/lastfm"
)

func getCustomArt(album string) string {
	fmt.Println("getting custom art")

	db, err := sql.Open("sqlite3", "./art.db")
	defer db.Close()

	checkError(err, "db opened")
	escapedAlbum := strings.ReplaceAll(album, "'", "''")
	query := "select arturl from artkv where album = '" + escapedAlbum + "'"

	fmt.Println("Art query: ", query)
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		fmt.Println("error on query")

	}
	var art string
	for rows.Next() {
		err = rows.Scan(&art)
		checkError(err, "scanning rows")
	}

	checkError(err, "getting custom art")

	return art

}

func getAlbumArtAndPlayCount(artist string, track string) (string, string) {
	// refactor this out!
	apikey := apiKey
	apisecret := apiPassword

	api := lastfm.New(apikey, apisecret)

	// Make this checking better
	trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": track})
	userResult, _ := api.User.GetRecentTracks(lastfm.P{"user": "dbrowning"})
	mbid := trackInfo.Album.Mbid
	fmt.Print("mbid: ", mbid)

	playcount := trackInfo.PlayCount

	// check for custom art first, if we have that then bail
	if len(userResult.Tracks) > 0 {
		customArt := getCustomArt(userResult.Tracks[0].Album.Name)

		if len(customArt) > 0 {
			artPath := "static/art/" + customArt
			return artPath, playcount
		}
	}

	images := trackInfo.Album.Images
	if len(images) == 0 {
		newTrackName := removeRemastered(track)
		trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": newTrackName})
		fmt.Println("NEW TRACK NAME: ", newTrackName)
		if len(images) == 0 {
			return "static/art/placeholder.jpg", playcount
		} else {
			images = trackInfo.Album.Images
			return images[2].Url, playcount
		}

	}

	fmt.Println("Images: ", images)

	largeImage := images[2].Url
	return largeImage, playcount
}
