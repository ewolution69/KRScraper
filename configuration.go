package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Configuration struct {
	ApiKey string
}

// load the configurationfile
func loadConfigurationFromFile() Configuration {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) // get path from executable
	file, _ := os.Open(dir + "/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	file.Close()
	return configuration
}
