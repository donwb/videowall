package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getArtistStats(artist string) *AristStats {
	connectString := os.Getenv("CONN")
	fmt.Println("Connection: ", connectString)

	if len(connectString) == 0 {
		return nil
	}
	fmt.Println(connectString)

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
