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

func getTopArtistsForDate(dateToCheck string, numToReturn int) []TopArtists {
	connectString := getConnection()

	db, err := sql.Open("postgres", connectString)
	defer db.Close()

	checkError(err, "opening connection")

	topQuery := `select artist, count(artist) as total
	from history
	where utc_time like $1
	group by artist order by total desc;`

	rows, err := db.Query(topQuery, dateToCheck)
	checkError(err, "Failed running top query")
	defer rows.Close()

	retVal := make([]TopArtists, 0)
	loopCounter := 0

	for rows.Next() {
		var art string
		var total int

		err := rows.Scan(&art, &total)
		checkError(err, "Error on scan")

		a := TopArtists{
			Artist:    art,
			Playcount: total,
		}

		retVal = append(retVal, a)
		loopCounter++
		if loopCounter >= numToReturn {
			break
		}

	}

	return retVal
}

func getConnection() string {
	connectString := os.Getenv("CONN")

	if len(connectString) == 0 {
		fmt.Println("the connection string is empty!")
		return ""
	}

	return connectString
}
