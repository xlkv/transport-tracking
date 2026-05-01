package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *App) Router() http.Handler {
	r := chi.NewRouter()

	return r
}
