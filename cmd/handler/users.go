package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nullcarter/nossle/cmd/services"
	"github.com/nullcarter/nossle/internal/store"
)

type UserHandler struct {
	Services services.Services
}

func (uh UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := uh.Services.Users.GetUsers(r.Context())

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error","Failed to fetch users")
		return
	}

	uh.Services.Response.Success(w, 200, users)

}
func (uh UserHandler) Get(w http.ResponseWriter, r *http.Request) {

}
func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		uh.Services.Error(w, 400, "internal_error", "Failed to create user.")
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		validationError := err.(validator.ValidationErrors)
		uh.Services.Response.Error(w, http.StatusBadRequest, "internal_error", validationError.Error())
		return
	}

	defer r.Body.Close()

	userParams := store.CreateUserParams {
		Username: user.Username,
		PwHash: sql.NullString{String: user.Password, Valid: true},
	}

	err := uh.Services.Users.CreateUser(userParams, r.Context())

	if err != nil {
		uh.Services.Error(w, 400, "internal_error", "Failed to create user.")
		return
	}

	uh.Services.Success(w, 200, nil)
}
func (uh UserHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (uh UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
