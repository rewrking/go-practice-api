package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/charmbracelet/log"
)

const API_VERSION uint32 = 1

type JsonResultInfo struct {
	ApiVersion uint32 `json:"apiVersion"`
	Method     string `json:"method"`
}

type JsonErrorMessage struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type JsonError struct {
	JsonResultInfo
	Error JsonErrorMessage `json:"error"`
}

type JsonDataItems[T any] struct {
	Items T `json:"items"`
}

type JsonDataSlice[T any] struct {
	JsonResultInfo
	Data JsonDataItems[T] `json:"data"`
}

type JsonData[T any] struct {
	JsonResultInfo
	Data T `json:"data"`
}

func getDataItemsStruct[T any](data T, r *http.Request) JsonDataSlice[T] {
	return JsonDataSlice[T]{
		JsonResultInfo: JsonResultInfo{
			ApiVersion: API_VERSION,
			Method:     r.Method,
		},
		Data: JsonDataItems[T]{
			Items: data,
		},
	}
}

func getDataStruct[T any](data T, r *http.Request) JsonData[T] {
	return JsonData[T]{
		JsonResultInfo: JsonResultInfo{
			ApiVersion: API_VERSION,
			Method:     r.Method,
		},
		Data: data,
	}
}

func getErrorStruct(code uint32, r *http.Request) (JsonError, int) {
	var message string
	switch code {
	case http.StatusNotFound:
		message = "Resource Not Found"
	case http.StatusBadRequest:
		message = "Bad request"
	default:
		message = "Internal Server Error"
	}
	return JsonError{
		JsonResultInfo: JsonResultInfo{
			ApiVersion: API_VERSION,
			Method:     r.Method,
		},
		Error: JsonErrorMessage{
			Code:    code,
			Message: message,
		},
	}, int(code)
}

func setDefaultHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func writeNotFound(w http.ResponseWriter, r *http.Request) {
	result, code := getErrorStruct(http.StatusNotFound, r)
	res, _ := json.Marshal(result)
	log.Error(fmt.Sprint("[", code, "] ", r.URL), "err", result.Error.Message)
	setDefaultHeader(w)
	w.WriteHeader(code)
	w.Write(res)
}

func writeBadRequest(w http.ResponseWriter, r *http.Request) {
	result, code := getErrorStruct(http.StatusBadRequest, r)
	res, _ := json.Marshal(result)
	log.Error(fmt.Sprint("[", code, "] ", r.URL), "err", result.Error.Message)
	setDefaultHeader(w)
	w.WriteHeader(code)
	w.Write(res)
}

func writeDefaultHeader[T any](w http.ResponseWriter, r *http.Request, data T) {
	var result any
	code := http.StatusOK
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		result = getDataItemsStruct[T](data, r)
	} else {
		result = getDataStruct[T](data, r)
	}
	res, _ := json.Marshal(result)
	log.Info(fmt.Sprint("[", code, "] [", r.Method, "] ", r.URL))
	setDefaultHeader(w)
	w.WriteHeader(code)
	w.Write(res)
}
