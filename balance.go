package main

import (
	"math/rand"
	"net/http"
)

func balanceHander(handlers ...http.Handler) http.Handler {
	if len(handlers) == 1 {
		return handlers[0]
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers[rand.Intn(len(handlers))].ServeHTTP(w, r)
	})
}
