package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/atlokeshsk/snippetbox/internal/models"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	snippets     *models.SnippetModel
	templaeCache map[string]*template.Template
	formDecoder  *form.Decoder
}

func main() {
	// Command line arguments
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "Mysql datasource name")
	flag.Parse()

	// CustomLoggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err.Error())
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// application contains the dependcies
	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		snippets:     &models.SnippetModel{DB: db},
		templaeCache: templateCache,
		formDecoder:  form.NewDecoder(),
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("starting server in port :4000")
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
