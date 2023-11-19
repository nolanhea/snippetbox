package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errLog  *log.Logger
	infoLog *log.Logger
}

func main() {

	cfg := config{}
	flag.StringVar(&cfg.addr, "addr", ":4000", "Port to listen to")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Directory containing static variables")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errLog:  errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Println("Starting server on :80")
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
