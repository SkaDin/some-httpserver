package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"some-httpserver/internal/app/models"
	"some-httpserver/internal/app/processor"
	"strconv"
	"strings"
)

type UserHandler struct {
	processor *processor.UsersProcessor
}

func NewUsersHandler(processor *processor.UsersProcessor) *UserHandler {
	handler := &UserHandler{
		processor: processor,
	}
	return handler
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		WrapError(w, err)

		return
	}

	err = handler.processor.CreateUser(newUser)

	if err != nil {
		WrapError(w, err)

		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   "",
	}
	WrapOK(w, m)
}

func (handler *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	list, err := handler.processor.ListUsers(strings.Trim(vars.Get("name"), "\""))

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]any{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)

}

func (handler *UserHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	user, err := handler.processor.FindUser(id)
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]any{
		"result": "OK",
		"data":   user,
	}
	WrapOK(w, m)
}
