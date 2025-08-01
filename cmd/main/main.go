package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/nullcarter/nossle/cmd/api"
	"github.com/nullcarter/nossle/internal/store"
)

func main() {
	cfg := api.Config{
		Addr:   ":8080",
		DbPath: "sqlite",
	}

	db, err := sql.Open("sqlite", cfg.DbPath)

	if err != nil {
		log.Fatal(err)
	}

	appStore := store.NewStore(db)

	nossle := &api.Nossle{
		Config: cfg,
		Store:  appStore,
	}

	mux := nossle.Mount()

	log.Fatal(nossle.Run(mux))
}
