package add

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler for adding new item in bloom filter
func Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item := params["item"]
	if err := add(item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err := json.NewEncoder(w).Encode("Item added"); err != nil {
		http.Error(w, "Error encoding result to json: "+err.Error(), http.StatusInternalServerError)
	}
}

// It will add item in bloom filter
func add(i string) error {
	return nil
}
