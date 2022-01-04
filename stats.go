package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getArtistStats(artist string) *AristStats {

	connectString := getConnection()

	db, err := sql.Open("postgres", connectString)
	defer db.Close()

	checkError(err, "opening connection")

	rankQuery := `SELECT * FROM artist_ranking where artist = $1`
	row := db.QueryRow(rankQuery, artist)

	var ranking int
	var art string
	var total int
	row.Scan(&ranking, &art, &total)

	stats := &AristStats{
		ranking:   ranking,
		playCount: total,
	}
	totalRankQuery := `select max(ranking) as total_rank from artist_ranking`
	row = db.QueryRow(totalRankQuery)
	var totalRanks int
	row.Scan(&totalRanks)
	stats.maxRanks = totalRanks

	return stats
}

func getTrackStats(trackName string) *TrackStats {
	// Get Track stats
	connectString := getConnection()
	db, err := sql.Open("postgres", connectString)
	checkError(err, "opening connection")
	defer db.Close()

	trackPlayQuery := `select count(*) from history where track = $1`
	row := db.QueryRow(trackPlayQuery, trackName)

	var trackCount int
	row.Scan(&trackCount)
	trackStats := &TrackStats{
		playcount: trackCount,
	}

	return trackStats
}

func getAlbumStats(albumName string) *AlbumStats {
	connectString := getConnection()
	db, err := sql.Open("postgres", connectString)
	checkError(err, "opening connection")
	defer db.Close()

	albumQuery := `select count(*) from history where album = $1`
	row := db.QueryRow(albumQuery, albumName)

	var albumCount int
	row.Scan(&albumCount)
	albumStats := &AlbumStats{
		playcount: albumCount,
	}

	return albumStats
}

func getConnection() string {
	connectString := os.Getenv("CONN")
	fmt.Println("Connection: ", connectString)

	if len(connectString) == 0 {
		return ""
	}

	fmt.Println(connectString)
	return connectString
}
