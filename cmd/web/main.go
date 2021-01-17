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

	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func newApp() *app {
	a := app{}

	a.conf = startupConfig{}
	a.mux = http.NewServeMux()

	return &a
}

func (a *app) run() {
	a.conf.setup()

	a.infoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	a.errorLogger = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	a.configureRoutes()

	srv := &http.Server{
		Addr:     a.conf.Address,
		ErrorLog: a.errorLogger,
		Handler:  a.mux,
	}

	a.infoLogger.Printf("Starting server on %s", a.conf.Address)

	a.infoLogger.Fatal(srv.ListenAndServe())
}
