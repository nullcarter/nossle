package store

import (
	"context"
	"database/sql"
	"log"
)

type Store struct {
	User interface {
		Initialize() error
		Create(context.Context) error
		Get(context.Context) error
		Update(context.Context) error
		Delete(context.Context) error
	}
}

func NewStore(db *sql.DB) Store {
	userStore := &UserStore{db}
	err := userStore.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	return Store{
		User: userStore,
	}
}
