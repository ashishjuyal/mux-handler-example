package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandlers struct {
}

func (h CustomerHandlers) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/customers", h.getAllCustomers)
}

func (h CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	msg := map[string]string{"msg": "This route is defined inside private method"}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

}
