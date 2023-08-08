package routes

import (
	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/controllers"
)

func MakeCrudRoutes[T any, PT controllers.BMPtr[T]](router *mux.Router, route string, ctrlr *controllers.ModelCtrlr[T, PT]) bool {
	if route[0] != '/' {
		return false
	}

	routeWithId := route + "/{id}"

	router.HandleFunc(route, ctrlr.CreateOne).Methods("POST")
	router.HandleFunc(route, ctrlr.ReadAll).Methods("GET")
	router.HandleFunc(routeWithId, ctrlr.ReadOne).Methods("GET")
	router.HandleFunc(routeWithId, ctrlr.UpdateOne).Methods("PUT")
	router.HandleFunc(routeWithId, ctrlr.DeleteOne).Methods("DELETE")

	return true
}
