package global

import (
	"encoding/json"
	"net/http"
)

func HeartBeatHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  "success",
		"message": "Server is running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
