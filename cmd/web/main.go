package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	newApp().run()
}

type app struct {
	cfg      appConfig
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *pgxpool.Pool
}

func newApp() *app {
	a := new(app)

	a.cfg = appConfig{}
	a.cfg.setup()

	return a
}

func (a *app) addPgxpool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), a.cfg.connString)

	if err != nil {
		a.errorLog.Fatal(err.Error())
	}

	a.db = pool

	return pool
}

func (a *app) run() {
	a.cfg.setup()

	a.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	a.errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool := a.addPgxpool()
	defer pool.Close()

	srv := &http.Server{
		Addr:     a.cfg.addr,
		ErrorLog: a.errorLog,
		Handler:  a.routes(),
	}

	a.infoLog.Printf("Starting server on %s", a.cfg.addr)

	a.infoLog.Fatal(srv.ListenAndServe())
}
