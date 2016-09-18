package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "./proxa.toml", "path to config file")
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

	serverHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("X-Request-ID", uuid.NewV4().String())

		for _, route := range server.Routes {
			if strings.HasPrefix(r.URL.Path, route.Prefix) {
				backendHandlers[route.Backend].ServeHTTP(w, r)
				return
			}
		}

		http.NotFound(w, r)
	})

	log.Printf("Starting Serer %+v\n", server)
	err = http.ListenAndServe(server.Bind, logMiddleware(serverHandler))

	if err != nil {
		log.Fatalln(err.Error())
	}
}
