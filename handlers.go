package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	//"github.com/labstack/echo/v4/middleware"
	"github.com/shkh/lastfm-go/lastfm"
)

func renderWall(c echo.Context) error {
	apikey := apiKey
	apisecret := apiPassword

	fmt.Println(apiKey + "   " + apiPassword)

	api := lastfm.New(apikey, apisecret)

	userResult, _ := api.User.GetRecentTracks(lastfm.P{"user": "dbrowning"})

	isPlaying := isNowPlaying(userResult)

	if isPlaying {
		artist := userResult.Tracks[0].Artist.Name
		album := userResult.Tracks[0].Album
		track := userResult.Tracks[0].Name
		albumName := album.Name

		info := "is now playing"
		//result = artist + " " + album.Name + " " + info

		fmt.Println(artist + " " + albumName + " " + info)

		// Get album art
		albumArtURL, playcount := getAlbumArtAndPlayCount(artist, track)

		fmt.Println("URL: ", albumArtURL)

		artistStats := getArtistStats(artist)
		trackStats := getTrackStats(track)
		albumStats := getAlbumStats(albumName)

		// See if we should mask the artist or album name
		safeName := pmrcLookup(artist)
		if len(safeName) > 0 {
			artist = safeName
		}
		safeAlbum := pmrcLookup(albumName)
		if len(safeAlbum) > 0 {
			albumName = safeAlbum
		}
		trackName := pmrcLookup(track)
		if len(trackName) > 0 {
			track = trackName
		}

		pi := &PlayingInfo{
			Artist:            artist,
			Album:             albumName,
			Track:             track,
			Art:               albumArtURL,
			TotalPlayCount:    playcount,
			ArtistRanking:     artistStats.ranking,
			MyArtistPlayCount: artistStats.playCount,
			MaxRanks:          artistStats.maxRanks,
			MyTrackPlayCount:  trackStats.playcount,
			MyAlbumPlayCount:  albumStats.playcount,
		}
		res := &VideoWallResult{
			Idle:       false,
			NowPlaying: pi,
		}
		fmt.Println("")
		return c.JSONPretty(200, res, " ")

	} else {
		res := &VideoWallResult{
			Idle: true,
		}
		fmt.Println("")
		return c.JSONPretty(200, res, " ")
	}

}

func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "hdtv.html", map[string]interface{}{
		"name": "Don",
		"msg":  "Last FM",
	})
}

func squareHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Don",
		"msg":  "Last FM",
	})
}

func tvHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "tv.html", map[string]interface{}{
		"name": "Don",
		"msg":  "Last FM",
	})
}

func hdlowHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "hdtvlow.html", map[string]interface{}{
		"name": "Don",
		"msg":  "Last FM",
	})
}

func testHandler(c echo.Context) error {

	artistStats := getArtistStats("Public Enemy")

	ret := strconv.Itoa(artistStats.playCount)
	r := strconv.Itoa(artistStats.ranking)
	return c.String(200, ret+" ======= "+r)
}

func statsHandler(c echo.Context) error {

	rawDate := c.Param("date")
	rawTotal := c.Param("total")

	// yeah passing the Sql query, i know, it's janky but whatevs
	date := "%" + rawDate + "%"
	total, err := strconv.Atoi(rawTotal)
	if err != nil {
		total = 5
	}

	topStats := getTopArtistsForDate(date, total)

	return c.JSONPretty(200, topStats, " ")
}
