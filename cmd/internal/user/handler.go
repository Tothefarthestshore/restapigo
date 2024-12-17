package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapitry/cmd/internal/handlers"
)

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/usersURL", h.Getlist)
	router.POST("/usersURL", h.CreateUser)
	router.GET("/userURL", h.GetUserByUUID)
	router.PUT("/userURL", h.UpdateUser)
	router.PATCH("/userURL", h.PartiallyUpdateUser)
	router.DELETE("/userURL", h.DeleteUser)
}

func (h *handler) Getlist(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is user by uuid"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is delete user"))
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list partially update user"))
}
