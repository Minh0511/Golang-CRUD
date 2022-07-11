package Update

import "database/sql"

// UpdateRating update database
func NewMovieRating(db *sql.DB, movieName string, newRating float32) {
	_, err := db.Exec("UPDATE GoodMovies SET Rating = ? WHERE MovieName = ?", newRating, movieName)
	if err != nil {
		panic(err.Error())
	}
}
