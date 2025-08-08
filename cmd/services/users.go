package services

import (
	"context"

	"github.com/nullcarter/nossle/internal/store"
)

type Users struct {
	Store *store.Queries
}

func (u Users) GetUsers(ctx context.Context) ([]store.GetUsersRow, error) {
	users, err := u.Store.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
