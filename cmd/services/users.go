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

func (u Users) CreateUser(userParams store.CreateUserParams, ctx context.Context) error {
	err := u.Store.CreateUser(ctx, userParams)

	if err != nil {
		return err
	}

	return nil
}

func (u Users) GetUser(userId int64, ctx context.Context) (store.GetUserRow, error) {
	user, err := u.Store.GetUser(ctx, userId)

	if err != nil {
		return store.GetUserRow{}, err
	}

	return user, nil
}
