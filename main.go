package main

import (
	"log"
	"net/http"
	"todo/controllers"
	"todo/middleware"
	"todo/store"

	"github.com/gorilla/mux"
)

func main() {
	store.CreateList()

	router := mux.NewRouter()

	// Static Files
	fileServer := http.FileServer(http.Dir("static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Frontend Routes
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/faq", controllers.Faq)
	router.HandleFunc("/contact", controllers.Contact)

	// API
	router.HandleFunc("/api/create", controllers.CreateItem)
	router.HandleFunc("/api/delete", controllers.DeleteItem)

	log.Println("Server listening on port 3000.")
	http.ListenAndServe(":3000", middleware.TrimPath(router))
}
