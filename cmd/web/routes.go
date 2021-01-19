package main

import "net/http"

func (a *app) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(a.cfg.staticFiles))))

	return mux
}
