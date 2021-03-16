package main

import (
	"path/filepath"
	"text/template"
	"time"

	"github.com/curtisvermeeren/snippetbox/pkg/forms"
	"github.com/curtisvermeeren/snippetbox/pkg/models"
)

// templateData acts as a holding structure for dta passsed to HTML templates.
type templateData struct {
	CurrentYear int
	Form        *forms.Form
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// Returns a nicely formatted date string
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a FuncMap
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache creates a new in-memory map of templates
// A cache in memory avoids reading from the disk everytime a template is rendered
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		templateSet, err = templateSet.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet

	}

	return cache, nil
}
