package main

import (
	"github.com/ashishjuyal/mux-handler-examples/approach-one/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	ch := handlers.CustomerHandlers{}

	router := mux.NewRouter()

	// register all your routes here...

	/**
		Here all the routes are registered inside the main function. In this case the routes
		need to be defined as exported functions inside the handlers package. This will work
		but will also export the GetAllCustomers method making it accessible from anywhere
		(producing a code smell).
	**/
	router.HandleFunc("/customers", ch.GetAllCustomers)

	// starting mux server from main
	log.Fatal(http.ListenAndServe(":8000", router))
}
