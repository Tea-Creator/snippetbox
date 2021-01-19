package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	newApp().run()
}

type startupConfig struct {
	Address     string
	StaticFiles string
	ConnString  string
}

func (conf *startupConfig) setup() {
	flag.StringVar(&conf.Address, "addr", ":4000", "HTTP network address")
	flag.StringVar(&conf.StaticFiles, "static_files", "./ui/static/", "Directory with static content")
	flag.StringVar(&conf.ConnString, "conn_string", "", "Database Connection String")
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

	pool, err := pgxpool.Connect(context.Background(), a.conf.ConnString)

	if err != nil {
		a.errorLogger.Fatal(err.Error())
	}

	defer pool.Close()

	srv := &http.Server{
		Addr:     a.conf.Address,
		ErrorLog: a.errorLogger,
		Handler:  a.mux,
	}

	a.infoLogger.Printf("Starting server on %s", a.conf.Address)

	a.infoLogger.Fatal(srv.ListenAndServe())
}
