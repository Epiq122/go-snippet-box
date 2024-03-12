package main

import (
	"snippetbox.robertgleason.ca/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []*models.Snippet
}
