package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetViewById)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreateForm)
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeader)
	return standard.Then(router)
}
