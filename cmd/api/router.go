package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"tracking.xlkv.com/internal/handler"
	"tracking.xlkv.com/internal/repository"
	"tracking.xlkv.com/internal/service"
)

func (app *App) Router() http.Handler {
	r := chi.NewRouter()

	// auth & user
	userRepo := repository.NewUserRepository(app.DB.Pool)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(app.rs, app.cfg.JWT.SecretKey)
	userHandler := handler.NewUserHandler(*userService, *authService)

	// auth
	r.Group(func(r chi.Router) {
		r.Post("/api/register", userHandler.Register)
		r.Post("/api/login", userHandler.Login)
		r.Post("/api/refresh", userHandler.Refresh)
		r.Post("/api/logout", userHandler.Logout)
	})

	return r
}
