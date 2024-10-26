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

	mux.Get("/", app.ShowHome)
	mux.Get("/{page}", app.ShowPage)
	return mux
}
