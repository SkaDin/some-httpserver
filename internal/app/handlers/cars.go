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

type CarsHandler struct {
	processor *processor.CarsProcessor
}

func NewCarsHandler(processor *processor.CarsProcessor) *CarsHandler {
	handler := &CarsHandler{
		processor: processor,
	}
	return handler
}

func (handler *CarsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newCar models.Car

	err := json.NewDecoder(r.Body).Decode(&newCar)

	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateCar(newCar)

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

func (handler *CarsHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	var userIdFilter uint64 = 0

	if vars.Get("userid") != "" {
		var err error
		userIdFilter, err = strconv.ParseUint(vars.Get("userid"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}
	list, err := handler.processor.ListCars(
		userIdFilter,
		strings.Trim(vars.Get("brand"), "\""),
		strings.Trim(vars.Get("colour"), "\""),
		strings.Trim(vars.Get("licence_plate"), "\""),
	)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   list,
	}

	WrapOK(w, m)
}

func (handler *CarsHandler) Find(w http.ResponseWriter, r *http.Request) {
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

	car, err := handler.processor.FindCar(id)

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]any{
		"result": "OK",
		"data":   car,
	}

	WrapOK(w, m)
}
