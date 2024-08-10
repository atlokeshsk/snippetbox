package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/view/{id}", snippetViewById)
	mux.HandleFunc(fmt.Sprintf("%s /snippet/create", http.MethodPost), snippetCreate)
	log.Println("starting server in port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
