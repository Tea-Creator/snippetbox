package main

import "net/http"

func (a *app) configureRoutes() {
	a.mux.HandleFunc("/", a.home)
	a.mux.HandleFunc("/snippet", a.showSnippet)
	a.mux.HandleFunc("/snippet/create", a.createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	a.mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
