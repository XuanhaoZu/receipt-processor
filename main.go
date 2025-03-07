package main

import (
	"net/http"
	"receipt-processor/handlers"
)

func main() {
	http.HandleFunc("/receipts/process", handlers.ProcessReceipt)
	http.HandleFunc("/receipts/", handlers.GetPoints)

	http.ListenAndServe(":8080", nil)
}
