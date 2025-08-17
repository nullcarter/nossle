package services

import (
	"context"
	"io"
	"net/http"

	"github.com/nullcarter/nossle/internal/store"
)

type Services struct {
	Response interface {
		Success(http.ResponseWriter, int, any)
		Error(http.ResponseWriter, int, string, string)
	}
	Validation interface {
		RequestBody(io.ReadCloser, any) error
	}
	Users interface {
		GetUsers(context.Context) ([]store.GetUsersRow, error)
		CreateUser(store.CreateUserParams, context.Context) error
		GetUser(int64, context.Context) (store.GetUserRow, error)
		UpdateUser(int64, store.UpdateUserParams, context.Context) error
	}
}

func NewService(store *store.Queries) Services {
	return Services{
		Response:   Response{},
		Validation: Validation{},
		Users:      Users{store},
	}
}
