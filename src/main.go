package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/models"
	"github.com/rewrking/go-practice-api/pkg/routes"
	"gorm.io/driver/sqlite"
)

func main() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic("Failed to set time zone")
	}
	time.Local = loc // -> this is setting the global timezone

	router := mux.NewRouter()
	models.Initialize(sqlite.Open("data.sqlite"))

	http.Handle("/", router)

	if routes.RegisterAll(router) {
		log.Printf("Server started on port 4000")
		log.Fatal(http.ListenAndServe("localhost:4000", router))
	} else {
		panic("Failed to register routes")
	}
}
