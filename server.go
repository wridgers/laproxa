package main

import (
	"github.com/BurntSushi/toml"
)

type serverConfiguration struct {
	Bind     string
	Routes   []route
	Backends []backend
}

type route struct {
	Prefix  string
	Backend string
}

func loadServerConfiguration(path string) (serverConfiguration, error) {
	var conf serverConfiguration

	_, err := toml.DecodeFile(path, &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
