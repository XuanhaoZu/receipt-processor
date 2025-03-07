package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/storage"
	"strings"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	if points, exists := storage.GetPoints(id); exists {
		response := map[string]int{"points": points}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(map[string]string{"error": "No receipt found for that ID."})

	}
}
