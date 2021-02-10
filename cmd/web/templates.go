package main

import "github.com/curtisvermeeren/snippetbox/pkg/models"

// templateData acts as a holding structure for dta passsed to HTML templates.
type templateData struct {
	Snippet *models.Snippet
}
