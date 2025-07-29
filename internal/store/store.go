package store

import (
	"context"
	"database/sql"
)

type Store struct {
	User interface {
		Create(context.Context) error
		Get(context.Context) error
		Update(context.Context) error
		Delete(context.Context) error
	}
}

func NewStore(db *sql.DB) Store {
	return Store{
		User: &UserStore{db},
	}
}
