package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type backend struct {
	Name  string
	Addrs []string
}

func addBackendHandler(m map[string]http.Handler, b backend) {
	var handlers []http.Handler

	for _, addr := range b.Addrs {
		target, _ := url.Parse("http://" + addr)
		handlers = append(handlers, httputil.NewSingleHostReverseProxy(target))
	}

	m[b.Name] = balanceHander(handlers...)
}
