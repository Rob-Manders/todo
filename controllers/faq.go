package controllers

import (
	"net/http"
	"todo/util"
	"todo/views"
)

var faqView *views.View

func Faq(w http.ResponseWriter, r *http.Request) {
	faqView = views.NewView("default", "views/pages/faq.gohtml")
	
	w.Header().Set("Content-Type", "text/html")
	util.Must(faqView.Render(w, nil))
}