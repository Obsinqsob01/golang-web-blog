package main

import (
	"fmt"
	"golang-web/actions"
	"golang-web/middlewares"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewares.Loggin(actions.HomeHandler()))

	port, err := os.LookupEnv("PORT")
	if !err {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Running server on port", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
