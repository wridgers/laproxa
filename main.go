package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "./proxa.json", "path to config file")
	flag.Parse()

	server, err := loadServerConfiguration(configFilePath)

	if err != nil {
		log.Fatalln(err.Error())
	}

	backendHandlers := make(map[string]http.Handler)

	for _, backend := range server.Backends {
		log.Printf("Loading Backend %+v\n", backend)
		addBackendHandler(backendHandlers, backend)
	}

	log.Printf("Loading Frontend %+v\n", server.Frontend)
	serverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range server.Frontend.Routes {
			if strings.HasPrefix(r.URL.Path, route.Prefix) {
				backendHandlers[route.Backend].ServeHTTP(w, r)
				return
			}
		}

		http.NotFound(w, r)
	})

	log.Printf("Starting Frontend %+v\n", server.Frontend)
	http.ListenAndServe(server.Frontend.Bind, logMiddleware(serverHandler))
}
