package Config

var MoviesName = []string{"Avengers", "The Eternals", "Overwatch", "Elden ring", "Dark souls",
	"The Last of Us", "Honkai impact", "FlyMe2theMoon", "Stardew Valley", "Doom"}
var MoviesGenre = []string{"Action", "Comedy", "Romance", "Horror", "Harem",
	"Isekai", "Sci-fi", "Gender bender", "Slice of life", "Fantasy"}
var MoviesDirector = []string{"Kevin Kaslana", "Raiden Mei", "Bronya Zaychik", "Seele", "Otto Apocalypse",
	"Murata Himeko", "Rita Rossweisse", "Hidetaka Miyazaki", "George Martin", "Michael Bay"}
var MoviesRating = []float64{9.1, 9.5, 9.3, 9.2, 8.7, 7.2, 5.4, 3.1, 9.9, 8.6}

type GoodMovies struct {
	MovieName  string  `db:"MovieName"`
	MovieGenre string  `db:"MovieGenre"`
	Director   string  `db:"Director"`
	Rating     float32 `db:"Rating"`
}
