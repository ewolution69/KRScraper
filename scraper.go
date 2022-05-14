package main

import (
	"flag"
	_ "fmt"
)

func main() {
	var movieName string
	flag.StringVar(&movieName, "m", "movie", "moviename to search")
	flag.Parse()

	requestMovieData(movieName)
}
