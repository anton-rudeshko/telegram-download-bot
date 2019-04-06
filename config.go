package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token   string `json:"bot_token"`
	Timeout int    `json:"bot_poll_timeout"` // optional
	Debug   bool   `json:"bot_debug"`        // optional

	Location       string   `json:"location"`
	AllowedUserIds []int    `json:"allowed_user_ids"`
	MimeWhitelist  []string `json:"mime_whitelist"` // optional
	SuccessText    string   `json:"success_text"`   // optional
}

func ReadConfig(filepath string) Config {
	log.Printf(`Reading config from "%s"`, filepath)

	fileData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf(`Could not read file "%s": %s`, filepath, err)
	}

	config := Config{
		Timeout:        60,
		Debug:          false,
		AllowedUserIds: []int{},
		MimeWhitelist:  []string{},
		SuccessText:    "\u2705 Done.",
	}
	err = json.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf(`Could not parse config "%s": %s`, fileData, err)
	}

	if config.Token == "" {
		log.Fatal("config.bot_token is required")
	}
	if config.Location == "" {
		log.Fatal("config.location is required")
	}

	log.Printf(`token is "%s"`, config.Token)

	return config
}
