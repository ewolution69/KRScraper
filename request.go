package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestMovieData(movieName string) {
	fmt.Println("Search for Moviedata: ", movieName)
	apiKey := loadConfigurationFromFile().ApiKey
	searchstring := "http://www.omdbapi.com/?t=" + movieName + "&apikey=" + apiKey

	resp, err := http.Get(searchstring)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	data := MovieData{}
	json.Unmarshal([]byte(body), &data)

	fmt.Println("Moviename: ", data.Title)
	fmt.Println("Release: ", data.Released)
	fmt.Println("Poster: ", data.Poster)
	fmt.Println("Genre: ", data.Genre)
	fmt.Println("IMDB ID: ", data.ImdbID)
}
