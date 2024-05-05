package main

import (
	"log"
	"net/http"
	h "goserver/handlers"
	j "goserver/jobs"
	d "goserver/database"
)

func main() {
	// Database
	go d.GetClient()

	// Static assets for client
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", http.StripPrefix("/public/", fileServer))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	// API endpoints
	http.HandleFunc("/api/create", h.CreateJob)
	http.HandleFunc("/api/jobs", h.QueryJob)
	http.HandleFunc("/api/pool", h.QueryPool)
	http.HandleFunc("/api/stop", h.StopJob)
	http.HandleFunc("/api/status", h.QueryStatus)

	http.HandleFunc("/api/examples", h.GetExamples)
	http.HandleFunc("/api/example/one", h.GetExample)
	http.HandleFunc("/api/example/update", h.UpdateExample)
	http.HandleFunc("/api/example/delete", h.DeleteExample)
	http.HandleFunc("/api/example/create", h.CreateExample)

	go j.JobLoop()

	// TLS
	port := ":443"
	log.Println("Listening on port", port)
	go log.Fatal(http.ListenAndServeTLS(port, "certificate.pem", "key.pem", nil))
}