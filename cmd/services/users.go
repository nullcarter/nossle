package services

import (
	"context"

	"github.com/nullcarter/nossle/internal/store"
)

type Users struct {
	store *store.Queries
}

func (u Users) GetUsers(ctx context.Context) ([]store.User, error) {
	return nil, nil
}
