package services

import (
	"context"

	"github.com/nullcarter/nossle/internal/store"
)

type Services struct {
	Response
	Users interface {
		GetUsers(context.Context) ([]store.GetUsersRow, error)
	}
}

func NewService(store *store.Queries) Services {
	return Services{
		Response: Response{},
		Users: Users{store},
	}
}
