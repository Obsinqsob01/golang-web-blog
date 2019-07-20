package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Loggin create log for each request
func Loggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

// HomeHandler home route
func HomeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}

func main() {
	mux := http.NewServeMux()

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
