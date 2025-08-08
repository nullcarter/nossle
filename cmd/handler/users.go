package handler

import (
	"net/http"

	"github.com/nullcarter/nossle/cmd/services"
)

type UserHandler struct {
	Services services.Services
}

func (uh UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := uh.Services.Users.GetUsers(r.Context())

	if err != nil {
		uh.Services.Response.Error(w, 400, "internal_error","Failed to fetch users")
	}

	uh.Services.Response.Success(w, 200, users)

}
func (uh UserHandler) Get(w http.ResponseWriter, r *http.Request)    {

}
func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {}
func (uh UserHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (uh UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
