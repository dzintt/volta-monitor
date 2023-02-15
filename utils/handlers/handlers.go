package handlers

import (
	"encoding/json"
	"log"
	"os"
	"volta-monitor/models"
)

var (
	Settings models.Settings
)

//auto runs
func init() {
	file, err := os.Open("settings.json")

	if err != nil {
		log.Fatal("ERROR: " + err.Error())
	}

	err = json.NewDecoder(file).Decode(&Settings)
}
