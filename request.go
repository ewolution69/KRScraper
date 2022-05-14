package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func requestData(movieName string) {
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
	sb := string(body)
	fmt.Println(sb)
}
