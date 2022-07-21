package main

import (
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

var apiKey string
var apiPassword string

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	apiKey = os.Getenv("APIKEY")
	apiPassword = os.Getenv("SECRET")
	fmt.Println("Starting server.....")
	fmt.Println("======= ENVVARS =======")
	fmt.Println("Key: " + apiKey + " Secret: " + apiPassword)
	fmt.Println("Connect String: ", os.Getenv("CONN"))
	fmt.Println("-------- ENVVARS ------------")

	e := echo.New()

	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	// Routes
	e.Static("/static", "view")
	e.GET("/", homeHandler)
	e.GET("/hdlow", hdlowHandler)
	e.GET("/tv", tvHandler)
	e.GET("/square", squareHandler)
	e.GET("/api", renderWall)
	e.GET("/test", testHandler)
	e.GET("/api/stats", statsHandler)

	// Start!
	e.Logger.Fatal(e.Start(":1323"))

}
