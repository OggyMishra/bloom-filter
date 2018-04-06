package hash

import (
	"encoding/json"
	"net/http"
)

// Handler for adding new item in bloom filter
func Handler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(getHashFuncCount()); err != nil {
		http.Error(w, "Error encoding result to json: "+err.Error(), http.StatusInternalServerError)
	}
}

func getHashFuncCount() int {
	return 2
}
