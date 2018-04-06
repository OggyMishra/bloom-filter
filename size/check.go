package size

import (
	"encoding/json"
	"net/http"
)

// Handler for adding new item in bloom filter
func Handler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(getBloomSize()); err != nil {
		http.Error(w, "Error encoding result to json: "+err.Error(), http.StatusInternalServerError)
	}
}

func getBloomSize() int {
	return 20
}
