package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Addr         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

type Application struct {
	Config Config
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(router chi.Router) {
		router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome!\n"))
		})
	})

	return r
}

func (app *Application) Run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: app.Config.WriteTimeout,
		ReadTimeout:  app.Config.ReadTimeout,
		IdleTimeout:  app.Config.IdleTimeout,
	}

	log.Printf("server has started at %s", app.Config.Addr)

	return srv.ListenAndServe()
}
