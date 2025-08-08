package services

import (
	"context"

	"github.com/nullcarter/nossle/internal/store"
)

type Services struct {
	Users interface {
		GetUsers(context.Context) ([]store.User, error)
	}
}

func NewService(store *store.Queries) Services {
	return Services{
		Users: Users{store},
	}
}
