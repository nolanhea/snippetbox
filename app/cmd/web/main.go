package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/nolanhea/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	addr      string
	staticDir string
}

type application struct {
	errLog        *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

type MySqlConfig struct {
	Username string
	Password string
	DbName   string
	Port     uint
	Host     string
}

func main() {

	cfg := config{}
	flag.StringVar(&cfg.addr, "addr", ":4000", "Port to listen to")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Directory containing static variables")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conf := MySqlConfig{
		Username: "root",
		Password: "rootpassword",
		DbName:   "snippetbox",
		Port:     3306,
		Host:     "database",
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		conf.Username, conf.Password,
		conf.Host, conf.Port, conf.DbName,
	)

	db, err := openDB(dsn + "?parseTime=true")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errLog:        errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	mux := app.routes()
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	infoLog.Println("Starting server")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

//go get -u github.com/go-sql-driver/mysql
//docker-compose exec database mariadb -uuser -ppassword
//show databases;
//use db;
