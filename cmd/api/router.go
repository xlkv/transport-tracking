package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"tracking.xlkv.com/internal/handler"
	"tracking.xlkv.com/internal/repository"
	"tracking.xlkv.com/internal/serivce"
)

func (app *App) Router() http.Handler {
	r := chi.NewRouter()

	statusRepo := repository.NewStatusRepository(app.DB.Pool)
	statusService := serivce.NewStatusService(statusRepo)
	statusHandler := handler.NewStatusHandler(statusService)

	r.Get("/status", statusHandler.Status)

	return r
}
