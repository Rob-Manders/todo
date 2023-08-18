package controllers

import (
	"net/http"
	"todo/store"
	"todo/util"
	"todo/views"
)

var homeView *views.View

func Home(w http.ResponseWriter, r *http.Request) {
	homeView = views.NewView("default", "views/pages/home.gohtml")

	items := store.ListInstance.Items

	w.Header().Add("Content-Type", "text/html")
	util.Must(homeView.Render(w, items))
}