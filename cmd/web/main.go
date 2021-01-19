package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Tea-Creator/snippetbox/pkg/models/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	newApp().run()
}

type app struct {
	cfg      appConfig
	infoLog  *log.Logger
	errorLog *log.Logger
	snippets *postgres.SnippetModel
}

func newApp() *app {
	a := new(app)

	a.cfg = appConfig{}
	a.cfg.setup()

	return a
}

func (a *app) run() {
	a.cfg.setup()

	a.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	a.errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := a.addPgxpool()

	if err != nil {
		a.errorLog.Fatalln(err.Error())
	}

	defer pool.Close()

	a.snippets = &postgres.SnippetModel{DB: pool}

	srv := &http.Server{
		Addr:     a.cfg.addr,
		ErrorLog: a.errorLog,
		Handler:  a.routes(),
	}

	a.infoLog.Printf("Starting server on %s", a.cfg.addr)

	a.infoLog.Fatal(srv.ListenAndServe())
}

func (a *app) addPgxpool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), a.cfg.connString)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
