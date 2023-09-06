package main

import (
	"log"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("calling auth middleware...")

		next.ServeHTTP(w, r)
	})
}
