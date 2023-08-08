package controllers

import (
	"encoding/json"
	"net/http"
)

func writeNotFound(w http.ResponseWriter) {
	data := make(map[string]string)
	data["message"] = "Resource Not Found"
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
}

func writeBadRequest(w http.ResponseWriter) {
	data := make(map[string]string)
	data["message"] = "Bad request"
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}

func writeDefaultHeader(w http.ResponseWriter, data any) {
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
