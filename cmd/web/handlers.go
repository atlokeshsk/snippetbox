package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/atlokeshsk/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
	// files := []string{
	// 	"./ui/html/pages/base.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// }
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displays the specific snippet"))
}

func (app *application) snippetViewById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	title := "form handler"
	content := "from create snippet handler"
	expires := 365
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
