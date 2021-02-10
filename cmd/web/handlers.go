package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/curtisvermeeren/snippetbox/pkg/models"
)

// Handle home route
func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.notFound(w)
		return
	}

	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n\n", snippet)
	}

	/*
		files := []string{
			"../../ui/html/home.page.tmpl",
			"../../ui/html/base.layout.tmpl",
			"../../ui/html/footer.partial.tmpl",
		}

		ts, err := template.ParseFiles(files...)

		if err != nil {
			app.serverError(w, err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			app.serverError(w, err)
			http.Error(w, "Internal Server Error", 500)
		}
	*/
}

// Show a snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippet: s}

	files := []string{
		"../../ui/html/show.page.tmpl",
		"../../ui/html/base.layout.tmpl",
		"../../ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

// Create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\n Climb the mountain\n but slowly, slowly\n\n - Snail"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}