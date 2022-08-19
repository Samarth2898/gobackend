package util

const (
	INR = "INR"
	USD = "USD"
	EUR = "EUR"
)

// isSupportedCurrency returns true if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, EUR:
		return true
	}
	return false
}
