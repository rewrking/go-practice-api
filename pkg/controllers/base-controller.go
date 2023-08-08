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

type BMPtr[T any] interface {
	Update(*T)
	*T
}

type ModelCtrlr[T any, PT BMPtr[T]] struct{}

func Make[T any, PT BMPtr[T]]() ModelCtrlr[T, PT] {
	ctrlr := ModelCtrlr[T, PT]{}
	return ctrlr
}

func (ctrlr ModelCtrlr[T, PT]) CreateOne(w http.ResponseWriter, r *http.Request) {
	var newItem T
	parseBody(r, &newItem)

	item := models.CreateOne[T](&newItem)
	writeDefaultHeader(w, item)
}

func (ctrlr ModelCtrlr[T, PT]) ReadAll(w http.ResponseWriter, r *http.Request) {
	items := models.ReadAll[T]()
	writeDefaultHeader(w, items)
}

func (ctrlr ModelCtrlr[T, PT]) ReadOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		writeBadRequest(w)
	} else {
		item := models.ReadOne[T](ID)
		if item == nil {
			writeNotFound(w)
		} else {
			writeDefaultHeader(w, item)
		}
	}
}

func (ctrlr ModelCtrlr[T, PT]) UpdateOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		writeBadRequest(w)
	} else {
		var updates T
		parseBody(r, &updates)

		item := models.ReadOne[T](ID)
		if item == nil {
			writeNotFound(w)
		} else {
			typedItem := (PT)(item)
			typedItem.Update(&updates)
			writeDefaultHeader(w, typedItem)
		}
	}
}

func (ctrlr ModelCtrlr[T, PT]) DeleteOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		writeBadRequest(w)
	} else {
		result := models.DeleteOne[T](ID)
		if result == nil {
			writeNotFound(w)
		} else {
			writeDefaultHeader(w, result)
		}
	}
}
