package handler

import (
	"net/http"

	"github.com/nullcarter/nossle/cmd/services"
)

type UserHandler struct {
	Services services.Services
}

func (uh UserHandler) List(w http.ResponseWriter, r *http.Request) {

}
func (uh UserHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (uh UserHandler) Create(w http.ResponseWriter, r *http.Request) {}
func (uh UserHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (uh UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
