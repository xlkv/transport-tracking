package main

import (
	"log"
	"log/slog"
	"net/http"

	"tracking.xlkv.com/internal/cache"
	"tracking.xlkv.com/internal/config"
	"tracking.xlkv.com/internal/database"
)

type App struct {
	DB  *database.DB
	rs  *cache.RedisCache
	cfg *config.Config
}

func NewApp() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		slog.Error("Env load failed", "err", err)
		return nil, err
	}

	db, err := database.New(&cfg.DB)

	if err != nil {
		slog.Error("DB initial error", "err", err)
		return nil, err
	}

	rs, err := cache.NewRedisCache(cfg.RedisUrl)

	if err != nil {
		slog.Error("Redis initial error", "err", err)
		return nil, err
	}

	app := &App{
		DB:  db,
		cfg: cfg,
		rs:  rs,
	}

	return app, nil
}

func (app *App) Run() {

	if err := http.ListenAndServe(":"+app.cfg.Server.Port, app.Router()); err != nil {
		log.Fatal(err)
	}
}
