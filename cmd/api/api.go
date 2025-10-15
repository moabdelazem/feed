package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moabdelazem/feed/internal/store"
)

type application struct {
	config config
	store  store.Store
}

type config struct {
	addr     string
	dbConfig dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "healthy"}`))
}

func (app *application) mount() *chi.Mux {
	v1Router := chi.NewRouter()

	v1Router.Use(middleware.Recoverer)
	v1Router.Use(middleware.RealIP)
	v1Router.Use(middleware.Logger)
	v1Router.Use(middleware.RequestID)

	// ? Each request handled by `v1Router` must complete within 60 seconds, or it will be automatically canceled.
	v1Router.Use(middleware.Timeout(time.Second * 60))

	v1Router.Get("/v1/health", app.checkHealthHandler)

	return v1Router
}

func (app *application) run(mux *chi.Mux) error {
	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	// TODO: Structured logging will be impelmented
	log.Printf("Our server started at http://localhost%s", app.config.addr)

	return server.ListenAndServe()
}
