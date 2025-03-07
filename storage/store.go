package storage

import (
	"receipt-processor/models"
	"receipt-processor/utils"
	"sync"
)

var (
	mu         sync.Mutex
	receipts   = make(map[string]models.Receipt)
	points     = make(map[string]int)
	breakdowns = make(map[string][]utils.PointBreakdown)
)

// SaveReceipt Save recepits & points & breakdown
func SaveReceipt(id string, receipt models.Receipt, pts int, breakdown []utils.PointBreakdown) {
	mu.Lock()
	defer mu.Unlock()
	receipts[id] = receipt
	points[id] = pts
	breakdowns[id] = breakdown
}

func GetPoints(id string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	pts, exists := points[id]
	return pts, exists
}

func GetBreakdown(id string) ([]utils.PointBreakdown, bool) {
	mu.Lock()
	defer mu.Unlock()
	bd, exists := breakdowns[id]
	return bd, exists
}

func GetReceipt(id string) (models.Receipt, bool) {
	mu.Lock()
	defer mu.Unlock()
	r, exists := receipts[id]
	return r, exists
}
