package Read

import (
	"Go_project/Config"
	"database/sql"
	"fmt"
)

// GetAllMovies get all movies from database
func GetAllMovies(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM GoodMovies")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var movies Config.GoodMovies
		err = rows.Scan(&movies.MovieName, &movies.MovieGenre, &movies.Director, &movies.Rating)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Movie name: ", movies.MovieName, "| Movie genre: ", movies.MovieGenre, "| Movie Director: ", movies.Director, "| Movie Rating: ", movies.Rating)
	}
}

// GetMovieByName select with condition
func GetMovieByName(db *sql.DB, movieName string) {
	rows, err := db.Query("SELECT * FROM GoodMovies WHERE MovieName = ? ORDER BY Rating DESC ", movieName)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var movies Config.GoodMovies
		err = rows.Scan(&movies.MovieName, &movies.MovieGenre, &movies.Director, &movies.Rating)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Movie name: ", movies.MovieName, "| Movie genre: ", movies.MovieGenre, "| Movie Director: ", movies.Director, "| Movie Rating: ", movies.Rating)
	}
}
