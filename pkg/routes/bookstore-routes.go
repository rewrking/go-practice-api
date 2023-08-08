package routes

import (
	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/controllers"
	"github.com/rewrking/go-practice-api/pkg/models"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	ctrlr := controllers.Make[models.Book]()

	router.HandleFunc("/books", ctrlr.Create).Methods("POST")
	router.HandleFunc("/books", ctrlr.GetAll).Methods("GET")
	router.HandleFunc("/books/{id}", ctrlr.GetById).Methods("GET")
	router.HandleFunc("/books/{id}", ctrlr.UpdateById).Methods("PUT")
	router.HandleFunc("/books/{id}", ctrlr.DeleteById).Methods("DELETE")
}
