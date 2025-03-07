package utils

import (
	"fmt"
	"math"
	"receipt-processor/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type PointBreakdown struct {
	Points int    `json:"points"`
	Reason string `json:"reason"`
}

func CalculatePoints(receipt models.Receipt) (int, []PointBreakdown) {
	points := 0
	var breakdown []PointBreakdown

	// One point for every alphanumeric character in the retailer name.
	retailPoints := 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			retailPoints++
		}
	}
	points += retailPoints
	breakdown = append(breakdown, PointBreakdown{retailPoints, fmt.Sprintf("%d points - retailer name has %d characters", retailPoints, retailPoints)})

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
		breakdown = append(breakdown, PointBreakdown{50, "50 points - total is a round dollar amount"})
	}

	// 25 points if the total is a multiple of 0.25
	if total/0.25 == float64(int(total/0.25)) {
		points += 25
		breakdown = append(breakdown, PointBreakdown{25, "25 points - total is a multiple of 0.25"})
	}

	// 5 points for every two items on the receipt
	itemPairs := len(receipt.Items) / 2
	itemPoints := itemPairs * 5
	points += itemPoints
	breakdown = append(breakdown, PointBreakdown{itemPoints, fmt.Sprintf("%d points - %d pairs(5 points each)", itemPoints, itemPairs)})

	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		descLen := len(strings.TrimSpace(item.ShortDescription))
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemBonus := int(math.Ceil(price * 0.2))
			points += itemBonus
			breakdown = append(breakdown, PointBreakdown{itemBonus, fmt.Sprintf("%d points - \"%s\" is %d characters (a multiple of 3), item price of %.2f * 0.2 = %.2f, rounded up is %d points",
				itemBonus, item.ShortDescription, descLen, price, price*0.2, itemBonus)})
		}
	}

	// 6 points if the day in the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
		breakdown = append(breakdown, PointBreakdown{6, "6 points - purchase day is odd"})
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
		breakdown = append(breakdown, PointBreakdown{10, "10 points - purchase time is between 14:00 and 16:00"})
	}
	fmt.Println("breakdown:", breakdown)

	return points, breakdown
}
