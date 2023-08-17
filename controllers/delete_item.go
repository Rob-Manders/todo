package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"
)

type Request struct {
	Title string `json:"title"`
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bytes, readError := io.ReadAll(r.Body)

	if readError != nil { w.WriteHeader(500) }

	var request Request
	decodeError := json.Unmarshal(bytes, &request)

	if decodeError != nil { w.WriteHeader(500) }

	error := models.DeleteItem(request.Title)

	if error != nil { w.WriteHeader(500) }
}