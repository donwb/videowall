package main

import "html/template"

type PlayingInfo struct {
	Artist         string `json:"artist"`
	Album          string `json:"album"`
	Track          string `json:"track"`
	Art            string `json:"art"`
	TotalPlayCount string `json:"totalPlayCount"`
}

type VideoWallResult struct {
	NowPlaying *PlayingInfo `json:"nowPlaying"`
	Idle       bool         `json:"idle"`
}

type TemplateRegistry struct {
	templates *template.Template
}
