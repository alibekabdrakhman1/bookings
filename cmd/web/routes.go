package main

import (
	"net/http"

	"github.com/alibekabdrakhman1/project/pkg/config"
	"github.com/alibekabdrakhman1/project/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix(("/static", fileSerfileServer)))

	return mux
}
