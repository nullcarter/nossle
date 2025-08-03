package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"

	"github.com/nullcarter/nossle/cmd/api"
	"github.com/nullcarter/nossle/internal/store"
)

func main() {
	cfg := api.Config{
		Addr:   ":8080",
		DbPath: "./sqlite",
	}

	db, err := sql.Open("sqlite", cfg.DbPath)

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		log.Fatal(err)
	}

	appStore := store.New(db)

	nossle := &api.Nossle{
		Config: cfg,
		Store:  appStore,
	}

	mux := nossle.Mount()

	log.Fatal(nossle.Run(mux))
}
