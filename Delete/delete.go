package Delete

import (
	"database/sql"
	"fmt"
	"time"
)

// DeleteAll delete all data from database
func DeleteAll(db *sql.DB) {
	start := time.Now()
	_, err := db.Exec("DELETE FROM GoodMovies")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}

// DeleteMovieFromDB delete movie with condition
func DeleteMovieFromDB(db *sql.DB, movieName string) {
	_, err := db.Exec("DELETE FROM GoodMovies WHERE MovieName = ?", movieName)
	if err != nil {
		panic(err.Error())
	}
}
