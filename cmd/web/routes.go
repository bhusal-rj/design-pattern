package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	//show the 404 route if the non of the routes get matched.
	mux.Use(middleware.Recoverer)

	//timeout the request if it takes more than 60 seconds to complete.
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	//test pages
	mux.Get("/test-patterns", app.TestPatterns)

	//factory routes
	mux.Get("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.Get("/api/cat-from-factory", app.CreateCatFromFactory)

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	return mux
}
