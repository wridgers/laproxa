package main

type frontendConfiguration struct {
	Bind   string          `json:"bind"`
	Routes []frontendRoute `json:"routes"`
}

type frontendRoute struct {
	Prefix  string `json:"prefix"`
	Backend string `json:"backend"`
}
