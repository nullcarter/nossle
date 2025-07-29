package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

func (us *UserStore) Create(ctx context.Context) error {
	return nil
}

func (us *UserStore) Get(ctx context.Context) error {
	return nil
}

func (us *UserStore) Update(ctx context.Context) error {
	return nil
}

func (us *UserStore) Delete(ctx context.Context) error {
	return nil
}
