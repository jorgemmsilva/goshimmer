package main

import (
	"encoding/json"
	"os"

	"github.com/iotaledger/goshimmer/client"
)

// config type that defines the config structure
type configuration struct {
	WebAPI         string           `json:"WebAPI,omitempty"`
	BasicAuth      client.BasicAuth `json:"basic_auth,omitempty"`
	ReuseAddresses bool             `json:"reuse_addresses"`
}

// internal variable that holds the config
var config = configuration{}

var configJSON = `{
	"WebAPI": "http://127.0.0.1:8080",
	"basic_auth": {
	  "enabled": false,
	  "username": "goshimmer",
	  "password": "goshimmer"
	},
	"reuse_addresses": false,
}`

// load the config file
func loadConfig() {
	// open config file
	file, err := os.Open("config.json")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}

		if err = os.WriteFile("config.json", []byte(configJSON), 0o644); err != nil {
			panic(err)
		}
		if file, err = os.Open("config.json"); err != nil {
			panic(err)
		}
	}
	defer file.Close()

	// decode config file
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	}
}
