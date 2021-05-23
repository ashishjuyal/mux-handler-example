package handlers

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

type CustomerHandlers struct {
}

func (h CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"msg": "This route is defined inside public method"}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

}
