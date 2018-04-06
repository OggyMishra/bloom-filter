package member

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler for adding new item in bloom filter
func Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	item := params["item"]
	if err := json.NewEncoder(w).Encode(isMember(item)); err != nil {
		http.Error(w, "Error encoding result to json: "+err.Error(), http.StatusInternalServerError)
	}
}

func isMember(item string) bool {
	return false
}
