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

func main() {

	cfg := config{}
	flag.StringVar(&cfg.addr, "addr", ":4000", "Port to listen to")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Directory containing static variables")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Println("Starting server on :80")
	err := http.ListenAndServe(cfg.addr, mux)
	errorLog.Fatal(err)
}
