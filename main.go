package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
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
		track := userResult.Tracks[0].Name

		info := "is now playing"
		//result = artist + " " + album.Name + " " + info

		fmt.Println(artist + " " + album.Name + " " + info)

		// Get album art
		albumArtURL := getAlbumArt(artist, track)
		fmt.Println("URL: ", albumArtURL)

		pi := &PlayingInfo{
			Artist: artist,
			Album:  album.Name,
			Track:  track,
			Art:    albumArtURL,
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
	Track  string `json:"track"`
	Art    string `json:"art"`
}

type VideoWallResult struct {
	NowPlaying *PlayingInfo `json:"nowPlaying"`
	Idle       bool         `json:"idle"`
}

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	fmt.Println("Starting main....")

	apiKey = os.Getenv("APIKEY")
	apiPassword = os.Getenv("SECRET")
	fmt.Println("======= ENVVARS =======")
	fmt.Println("Key: " + apiKey + " Secret: " + apiPassword)

	e := echo.New()

	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Static("/static", "view")
	e.GET("/", renderWall)
	e.GET("/home", HomeHandler)

	e.Logger.Fatal(e.Start(":1323"))

}

func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Don",
		"msg":  "Last FM",
	})
}

func getAlbumArt(artist string, track string) string {

	// refactor this out!
	apikey := apiKey
	apisecret := apiPassword

	api := lastfm.New(apikey, apisecret)

	// Make this checking better
	trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": track})
	images := trackInfo.Album.Images
	if len(images) == 0 {
		newTrackName := removeRemastered(track)
		trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": artist, "track": newTrackName})
		fmt.Println("NEW TRACK NAME: ", newTrackName)
		if len(images) == 0 {
			return "placeholder"
		} else {
			images = trackInfo.Album.Images
			return images[2].Url
		}

	}

	fmt.Println("Images: ", images)

	largeImage := images[2].Url
	return largeImage

	//return "https://lastfm.freetls.fastly.net/i/u/174s/e5078801aed03ec9bc933a736349f143.png"
}

func removeRemastered(trackName string) string {
	// try to remove the (remastered) from the end of a song if it gets in the way of album art
	return "Champagne Supernova"
}
