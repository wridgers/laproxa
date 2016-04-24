package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type backendConfiguration struct {
	Name  string   `json:"name"`
	Addrs []string `json:"addrs"`
}

func addBackendHandler(m map[string]http.Handler, b backendConfiguration) {
	var handlers []http.Handler

	for _, addr := range b.Addrs {
		target, _ := url.Parse("http://" + addr)
		handlers = append(handlers, httputil.NewSingleHostReverseProxy(target))
	}

	m[b.Name] = balanceHander(handlers...)
}
