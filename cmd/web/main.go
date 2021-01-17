package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// Port on witch server will start listening incoming requests
	Port = 4000
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on :%d\n", Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", Port), mux))
}
