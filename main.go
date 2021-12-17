package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	//"github.com/labstack/echo/v4/middleware"
	"github.com/shkh/lastfm-go/lastfm"
)

var apiKey string
var apiPassword string

func renderWall(c echo.Context) error {
	apikey := apiKey
	apisecret := apiPassword

	fmt.Println("here!")

	fmt.Println(apiKey + "   " + apiPassword)

	api := lastfm.New(apikey, apisecret)

	//var result string = ""
	fmt.Println("This is a test")

	userResult, _ := api.User.GetRecentTracks(lastfm.P{"user": "dbrowning"})

	isPlaying := isNowPlaying(userResult)

	if isPlaying {
		artist := userResult.Tracks[0].Artist.Name
		album := userResult.Tracks[0].Album
		info := "is now playing"
		//result = artist + " " + album.Name + " " + info

		fmt.Println(artist + " " + album.Name + " " + info)

		pi := &PlayingInfo{
			Artist: artist,
			Album:  album.Name,
		}
		res := &VideoWallResult{
			Idle:       false,
			NowPlaying: pi,
		}
		return c.JSONPretty(200, res, " ")

	} else {
		res := &VideoWallResult{
			Idle: true,
		}
		return c.JSONPretty(200, res, " ")
	}

}

type PlayingInfo struct {
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

type VideoWallResult struct {
	NowPlaying *PlayingInfo `json:"nowPlaying"`
	Idle       bool         `json:"idle"`
}

func main() {

	fmt.Println("Starting main....")

	apiKey = os.Getenv("APIKEY")
	apiPassword = os.Getenv("SECRET")

	fmt.Println("Key: " + apiKey + " Secret: " + apiPassword)

	e := echo.New()

	e.GET("/", renderWall)

	e.Logger.Fatal(e.Start(":1323"))

}
