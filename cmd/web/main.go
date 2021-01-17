package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	newApp().run()
}

type startupConfig struct {
	Address     string
	StaticFiles string
}

func (conf *startupConfig) setup() {
	flag.StringVar(&conf.Address, "addr", ":4000", "HTTP network address")
	flag.StringVar(&conf.StaticFiles, "static_files", "./ui/static/", "Directory with static content")
	flag.Parse()
}

type app struct {
	conf startupConfig
	mux  *http.ServeMux
}

func newApp() *app {
	a := app{}

	a.conf = startupConfig{}
	a.mux = http.NewServeMux()

	return &a
}

func (a *app) run() {
	a.conf = startupConfig{}
	a.conf.setup()

	a.configureHandlers()

	log.Printf("Starting server on %s\n", a.conf.Address)

	log.Fatal(http.ListenAndServe(*&a.conf.Address, a.mux))
}

func (a *app) configureHandlers() {
	a.mux.HandleFunc("/", home)
	a.mux.HandleFunc("/snippet", showSnippet)
	a.mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	a.mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
