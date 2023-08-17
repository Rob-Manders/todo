package controllers

import (
	"net/http"
	"todo/util"
	"todo/views"
)

var aboutView *views.View

func About(w http.ResponseWriter, r *http.Request) {
	aboutView = views.NewView("default", "views/pages/about.gohtml")
	
	w.Header().Set("Content-Type", "text/html")
	util.Must(aboutView.Render(w, nil))
}