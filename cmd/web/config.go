package main

import "flag"

type appConfig struct {
	addr        string
	staticFiles string
	connString  string
}

func (cfg *appConfig) setup() {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticFiles, "static_files", "./ui/static/", "Directory with static content")
	flag.StringVar(&cfg.connString, "conn_string", "", "Database Connection String")
	flag.Parse()
}
