package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nullcarter/nossle/cmd/handler"
	"github.com/nullcarter/nossle/internal/store"
)

type Config struct {
	Addr         string
	DbPath       string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

type Nossle struct {
	Config Config
	Store  *store.Queries
}

func (app *Nossle) Mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {

		userHandler := handler.UserHandler{}

		r.Route("/users", func(route chi.Router) {
			route.Get("/", userHandler.List)
			route.Post("/", userHandler.Create)
			route.Get("/{id}", userHandler.Get)
			route.Put("/{id}", userHandler.Update)
			route.Delete("/{id}", userHandler.Delete)
		})

	})

	return r
}

func (app *Nossle) Run(mux http.Handler) error {
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
