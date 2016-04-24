package main

import (
	"encoding/json"
	"io/ioutil"
)

type serverConfiguration struct {
	Frontend frontendConfiguration  `json:"frontend"`
	Backends []backendConfiguration `json:"backends"`
}

func loadServerConfiguration(configFilePath string) (serverConfiguration, error) {
	var conf serverConfiguration

	configData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(configData, &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
