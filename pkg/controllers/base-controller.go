package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rewrking/go-practice-api/pkg/models"
)

func parseBody[T any](r *http.Request, x *T) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func WriteNotFound(w http.ResponseWriter) {
	data := make(map[string]string)
	data["message"] = "Resource Not Found"
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(res)
}

func WriteBadRequest(w http.ResponseWriter) {
	data := make(map[string]string)
	data["message"] = "Bad request"
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}

func WriteDefaultHeader(w http.ResponseWriter, data any) {
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type BMPtr[T any] interface {
	Update(*T)
	*T
}

type ModelCtrlr[T any, PT BMPtr[T]] struct{}

func Make[T any, PT BMPtr[T]]() ModelCtrlr[T, PT] {
	ctrlr := ModelCtrlr[T, PT]{}
	return ctrlr
}

func (ctrlr ModelCtrlr[T, PT]) Create(w http.ResponseWriter, r *http.Request) {
	var newItem T
	parseBody(r, &newItem)

	item := models.Create[T](&newItem)
	WriteDefaultHeader(w, item)
}

func (ctrlr ModelCtrlr[T, PT]) GetAll(w http.ResponseWriter, r *http.Request) {
	items := models.GetAll[T]()
	WriteDefaultHeader(w, items)
}

func (ctrlr ModelCtrlr[T, PT]) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		WriteBadRequest(w)
	} else {
		item := models.GetById[T](ID)
		if item == nil {
			WriteNotFound(w)
		} else {
			WriteDefaultHeader(w, item)
		}
	}
}

func (ctrlr ModelCtrlr[T, PT]) UpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		WriteBadRequest(w)
	} else {
		var updates T
		parseBody(r, &updates)

		item := models.GetById[T](ID)
		if item == nil {
			WriteNotFound(w)
		} else {
			typedItem := (PT)(item)
			typedItem.Update(&updates)
			WriteDefaultHeader(w, typedItem)
		}
	}
}

func (ctrlr ModelCtrlr[T, PT]) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		WriteBadRequest(w)
	} else {
		result := models.DeleteById[T](ID)
		if result == nil {
			WriteNotFound(w)
		} else {
			WriteDefaultHeader(w, result)
		}
	}
}
