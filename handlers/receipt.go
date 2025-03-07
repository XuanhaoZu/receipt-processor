package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"

	"github.com/google/uuid"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		sendJSONError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	if err := utils.ValidateReceipt(receipt.Retailer, receipt.PurchaseDate, receipt.PurchaseTime, receipt.Total, len(receipt.Items)); err != nil {
		sendJSONError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points, breakdown := utils.CalculatePoints(receipt)
	storage.SaveReceipt(id, receipt, points, breakdown)

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
