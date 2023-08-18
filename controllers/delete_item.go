package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"todo/models"
)

type Request struct {
	Key string `json:"key"`
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	r.Body = http.MaxBytesReader(w, r.Body, 1 << 20)
	bytes, readError := io.ReadAll(r.Body)

	if readError != nil { w.WriteHeader(500) }

	var request Request
	decodeError := json.Unmarshal(bytes, &request)

	if decodeError != nil { w.WriteHeader(500) }

	error := models.DeleteItem(request.Key)

	if error != nil { w.WriteHeader(500) }
}