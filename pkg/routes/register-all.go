package routes

import "github.com/gorilla/mux"

func RegisterAll(router *mux.Router) bool {
	result := true
	result = result && RegisterBooks(router)
	return result
}
