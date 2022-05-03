package main

import (
	"log"
	"net/http"
	"playground-gin/app/handlers"
	"playground-gin/routes"
)

func main() {
	mux := http.NewServeMux()

	// static/public files
	fileServer := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public", fileServer))

	indexHandlers := handlers.IndexHandler{}
	routes.AddHandlers(mux, indexHandlers)

	log.Println("start web on port 80")
	err := http.ListenAndServe(":80", mux)
	log.Fatal(err)
}
