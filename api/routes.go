package api

import "github.com/gorilla/mux"

func CreateRoutes(userHandler *handlers.UserHandler, carsHandler *handlers.CarHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/create", userHandler.Create).Methods("POST")
	r.HandleFunc("/users/list", userHandler.List).Methods("GET")
	r.HandleFunc("/users/find/{id:[0-9]+}", userHandler.Find).Methods("GET")

	r.HandleFunc("/cars/create", carsHandler.Create).Methods("POST")
	r.HandleFunc("/cars/list", carsHandler.List).Methods("GET")
	r.HandleFunc("/cars/find/{id:[0-9]+}", carsHandler.Find).Methods("GET")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()

	return r
}
