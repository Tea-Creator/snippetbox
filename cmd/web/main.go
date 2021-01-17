package main

import (
	"flag"
	"log"
	"net/http"
	"os"
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
	a.conf.setup()

	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	a.configureHandlers()

	srv := &http.Server{
		Addr:     a.conf.Address,
		ErrorLog: errorLogger,
		Handler:  a.mux,
	}

	infoLogger.Printf("Starting server on %s", a.conf.Address)

	errorLogger.Fatal(srv.ListenAndServe())
}

func (a *app) configureHandlers() {
	a.mux.HandleFunc("/", home)
	a.mux.HandleFunc("/snippet", showSnippet)
	a.mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	a.mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
