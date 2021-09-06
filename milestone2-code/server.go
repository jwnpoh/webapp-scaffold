package main

import (
	"net/http"

	"github.com/jwnpoh/webapp-scaffold/milestone2-code/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/api", handlers.Api)
	mux.HandleFunc("/healthcheck", handlers.Healthcheck)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	s.ListenAndServe()
}
