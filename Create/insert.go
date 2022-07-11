package Create

import (
	"Go_project/Config"
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//insert data to database (slow version)
func InsertRandom(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO GoodMovies(MovieName, MovieGenre, Director, Rating) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_, err = stmt.Exec(Config.MoviesName[rand.Intn(len(Config.MoviesName))], Config.MoviesGenre[rand.Intn(len(Config.MoviesGenre))], Config.MoviesDirector[rand.Intn(len(Config.MoviesDirector))], Config.MoviesRating[rand.Intn(len(Config.MoviesRating))])
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Time elapsed: ", time.Since(start))
	if err != nil {
		panic(err.Error())
	}
}

//insert data to database (fast version)
func InsertBatch(db *sql.DB) {
	start := time.Now()
	for i := 1; i <= 10; i++ {
		generate := make([]interface{}, 0)
		placeholders := make([]string, 0)
		for j := 0; j < 1000; j++ {
			generate = append(generate, Config.MoviesName[rand.Intn(len(Config.MoviesName))],
				Config.MoviesGenre[rand.Intn(len(Config.MoviesGenre))], Config.MoviesDirector[rand.Intn(len(Config.MoviesDirector))],
				Config.MoviesRating[rand.Intn(len(Config.MoviesRating))])
			placeholders = append(placeholders, "(?, ?, ?, ?)")
		}
		movieInsert := fmt.Sprintf(`INSERT INTO GoodMovies(MovieName, MovieGenre, Director, Rating) VALUES %s
		`, strings.Join(placeholders, ","))

		_, err := db.Exec(movieInsert, generate...)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Time elapsed: ", time.Since(start))
}

//insert 1 data to database
func InsertMovies(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO GoodMovies(MovieName, MovieGenre, Director, Rating) VALUES(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec("Berserk", "Action", "Miura Kentaro", 9.1)
	if err != nil {
		panic(err.Error())
	}
}
