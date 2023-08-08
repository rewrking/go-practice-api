package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/config"
	"github.com/rewrking/go-practice-api/pkg/models"
	"github.com/rewrking/go-practice-api/pkg/routes"
	"gorm.io/driver/sqlite"
)

func main() {
	config.SetTimeZoneUTC()

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
