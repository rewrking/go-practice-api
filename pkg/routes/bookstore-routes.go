package routes

import (
	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/controllers"
	"github.com/rewrking/go-practice-api/pkg/models"
)

func RegisterBookStoreRoutes(router *mux.Router) bool {
	ctrlr := controllers.Make[models.Book]()
	return MakeCrudRoutes(router, "/books", &ctrlr)
}
