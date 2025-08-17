package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/nullcarter/nossle/cmd/services"
	"github.com/nullcarter/nossle/internal/store"
)

type UserHandler struct {
	Services services.Services
}

func (uh UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := uh.Services.Users.GetUsers(r.Context())

	// Services
	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to fetch users")
		return
	}

	uh.Services.Response.Success(w, 200, users)
}

func (uh UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to get user.")
		return
	}

	user, err := uh.Services.Users.GetUser(userId, r.Context())

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to get user.")
		return
	}
	uh.Services.Response.Success(w, 200, user)
}

func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Request Validation
	var user struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := uh.Services.Validation.RequestBody(r.Body, &user); err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", err.Error())
		return
	}

	// Services
	userParams := store.CreateUserParams{
		Username: user.Username,
		PwHash:   sql.NullString{String: user.Password, Valid: true},
	}

	err := uh.Services.Users.CreateUser(userParams, r.Context())

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to create user.")
		return
	}

	uh.Services.Response.Success(w, 200, nil)
}

func (uh UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Request Validation
	userIdStr := r.PathValue("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to update user.")
		return
	}

	// Services
	var user struct {
		Username string `json:"username"`
		RoleId string `json:"role_id"`
	}

	if err := uh.Services.Validation.RequestBody(r.Body, &user); err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to get user.")
		return
	}

	userUpdate := store.UpdateUserParams{
		ID: userId,
		Username: user.Username,
		// RoleID: user.RoleId,
	}

	if err := uh.Services.Users.UpdateUser(userId, userUpdate, r.Context()); err != nil {
		uh.Services.Response.Error(w, 400, "internal_error", "Failed to get user.")
		return
	}

	uh.Services.Response.Success(w, 200, nil)
}

func (uh UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
