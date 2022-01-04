package main

import "html/template"

type PlayingInfo struct {
	Artist            string `json:"artist"`
	Album             string `json:"album"`
	Track             string `json:"track"`
	Art               string `json:"art"`
	TotalPlayCount    string `json:"totalPlayCount"`
	ArtistRanking     int    `json:ranking`
	MyArtistPlayCount int    `json:myArtistsPlayCount`
	MaxRanks          int    `json:maxRanks`
	MyTrackPlayCount  int    `json:myTrackPlayCount`
	MyAlbumPlayCount  int    `json:myAlbumPlayCount`
}

type VideoWallResult struct {
	NowPlaying *PlayingInfo `json:"nowPlaying"`
	Idle       bool         `json:"idle"`
}

type TemplateRegistry struct {
	templates *template.Template
}

type AristStats struct {
	ranking   int
	playCount int
	maxRanks  int
}

type TrackStats struct {
	playcount int
}

type AlbumStats struct {
	playcount int
}
