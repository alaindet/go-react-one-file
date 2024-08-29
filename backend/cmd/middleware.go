package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)
  
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s %s\n", r.Method, r.URL.EscapedPath())
			next.ServeHTTP(w, r)
		},
	)
}

func corsMiddleware(next http.Handler) http.Handler {
	mw := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	return mw.Handler(next)
}