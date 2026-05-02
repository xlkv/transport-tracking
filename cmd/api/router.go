package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"tracking.xlkv.com/internal/handler"
	"tracking.xlkv.com/internal/middleware"
	"tracking.xlkv.com/internal/repository"
	"tracking.xlkv.com/internal/service"
	"tracking.xlkv.com/internal/ws"
)

func (app *App) Router() http.Handler {
	r := chi.NewRouter()

	userRepo := repository.NewUserRepository(app.DB.Pool)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(app.rs, app.cfg.JWT.SecretKey)
	userHandler := handler.NewUserHandler(*userService, *authService)

	authMiddleware := middleware.Auth(authService)

	locationRepo := repository.NewLocationRepository(app.DB.Pool)
	locationService := service.NewLocationService(locationRepo, app.rs)
	locationHandler := handler.NewLocationHandler(*locationService)

	vehicleWsHandler := ws.NewVehicleWSHandler(app.rs)

	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

		r.Post("/api/location", locationHandler.Create)
	})
	r.Group(func(r chi.Router) {
		r.Get("/ws/vehicle/{vehicleID}", vehicleWsHandler.HandleVehicleWS)
	})

	r.Group(func(r chi.Router) {
		r.Post("/api/register", userHandler.Register)
		r.Post("/api/login", userHandler.Login)
		r.Post("/api/refresh", userHandler.Refresh)
		r.Post("/api/logout", userHandler.Logout)
	})

	return r
}
