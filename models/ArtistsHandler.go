package models

import (
	"encoding/json"
	"net/http"

	"fiter/services"
)

func BandsHandler(w http.ResponseWriter, r *http.Request) {
	band, err := services.LoadBands("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(band)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
