package main

import (
	"log"

	"github.com/nullcarter/nossle/cmd/api"
)

func main() {
	cfg := api.Config{
		Addr: ":8080",
	}

	app := &api.Application{
		Config: cfg,
	}

	mux := app.Mount()
	log.Fatal(app.Run(mux))
}
