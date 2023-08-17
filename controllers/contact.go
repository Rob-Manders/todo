package controllers

import (
	"net/http"
	"todo/util"
	"todo/views"
)

var contactView *views.View

func Contact(w http.ResponseWriter, r *http.Request) {
	contactView = views.NewView("default", "views/pages/contact.gohtml")
	
	w.Header().Set("Content-Type", "text/html")
	util.Must(contactView.Render(w, nil))
}
