package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, err := json.Marshal(m)
	if err != nil {
		log.Println("Error Marshal JSON", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	if _, err = fmt.Fprintln(w, string(res)); err != nil {
		log.Println("Error conclusion JSON")
	}
}

func WrapOK(w http.ResponseWriter, m map[string]any) {
	res, err := json.Marshal(m)
	if err != nil {
		log.Println("Error Marshal JSON", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprintf(w, string(res)); err != nil {
		log.Println("Error conclusion JSON")
	}
}
