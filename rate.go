package bitpay

// A bbbRate represents a Bitcoin Best Bid (BBB) exchange rate
type bbbRate struct {
	Code string  `json:"code"`
	Name string  `json:"name"`
	Rate float64 `json:"rate"`
}
