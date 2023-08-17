package controllers

import (
	"net/http"
	"todo/util"
	"todo/views"
)

var homeView *views.View

func Home(w http.ResponseWriter, r *http.Request) {
	homeView = views.NewView("default", "views/pages/home.gohtml")

	w.Header().Set("Content-Type", "text/html")
	util.Must(homeView.Render(w, nil))
}