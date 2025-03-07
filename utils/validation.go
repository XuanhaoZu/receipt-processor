package utils

import (
	"errors"
	"regexp"
)

var (
	retailerRegex = regexp.MustCompile(`^[\w\s\-&]+$`)
	priceRegex    = regexp.MustCompile(`^\d+\.\d{2}$`)
	dateRegex     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	timeRegex     = regexp.MustCompile(`^\d{2}:\d{2}$`)
)

func ValidateReceipt(retailer, date, time, total string, items int) error {
	if !retailerRegex.MatchString(retailer) {
		return errors.New("Invalid retailer format")
	}
	if !dateRegex.MatchString(date) {
		return errors.New("Invalid purchaseDate format (YYYY-MM-DD)")
	}
	if !timeRegex.MatchString(time) {
		return errors.New("Invalid purchaseTime format (HH:MM)")
	}
	if !priceRegex.MatchString(total) {
		return errors.New("Invalid total format (e.g., 12.34)")
	}
	if items < 1 {
		return errors.New("At least one item is required")
	}
	return nil
}
