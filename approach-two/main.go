package main

import (
	"github.com/ashishjuyal/mux-handler-examples/approach-two/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	ch := handlers.CustomerHandlers{}

	router := mux.NewRouter()

	/**
		Here we are delegating the responsibility of registering the routes to the customer handler struct, that too inside the handlers package.
		In this case the routes don't need to be defined as exported functions and there visibility will remain inside handlers package.

		There is even a better third approach defined inside the https://github.com/ashishjuyal/banking application.
	**/
	// registering all your routes inside the customer handlers...
	ch.RegisterRoutes(router)

	// starting mux server from main
	log.Fatal(http.ListenAndServe(":8000", router))
}
