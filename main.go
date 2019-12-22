package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc("/real-time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Responding in real-time\n")
	})

	handler.HandleFunc("/1s", func(w http.ResponseWriter, r *http.Request) {
		<-time.After(time.Second)
		fmt.Fprintf(w, "Responding after 1 second\n")
	})

	handler.HandleFunc("/5s", func(w http.ResponseWriter, r *http.Request) {
		<-time.After(5 * time.Second)
		fmt.Fprintf(w, "Responding after 5 seconds\n")
	})

	handler.HandleFunc("/15s", func(w http.ResponseWriter, r *http.Request) {
		<-time.After(15 * time.Second)
		fmt.Fprintf(w, "Responding after 15 seconds\n")
	})

	handler.HandleFunc("/30s", func(w http.ResponseWriter, r *http.Request) {
		<-time.After(30 * time.Second)
		fmt.Fprintf(w, "Responding after 30 seconds\n")
	})

	// Config from https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handler,
		ReadTimeout:       1 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      90 * time.Second,
		IdleTimeout:       90 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
