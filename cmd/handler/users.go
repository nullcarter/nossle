package handler

import "net/http"

type UserHandler struct{}

func (b UserHandler) List(w http.ResponseWriter, r *http.Request)   {}
func (b UserHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (b UserHandler) Create(w http.ResponseWriter, r *http.Request) {}
func (b UserHandler) Update(w http.ResponseWriter, r *http.Request) {}
func (b UserHandler) Delete(w http.ResponseWriter, r *http.Request) {}
