package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"
	"todo/store"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bytes, readError := io.ReadAll(r.Body)

	if readError != nil {
		w.WriteHeader(500)
	}

	var item store.Item
	decodeError := json.Unmarshal(bytes, &item)

	if decodeError != nil {
		w.WriteHeader(500)
	}

	models.CreateItem(item)

	w.WriteHeader(201)
}
