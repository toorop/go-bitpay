package main

import (
	"fmt"
	"github.com/toorop/go-bitpay"
)

const (
	API_KEY = ""
)

func main() {
	bp := bitpay.New(API_KEY)

	// Get Bitcoin Best Bid Rates
	/*bbbr, err := bp.GetBitcoinBestBidRates()
	fmt.Println(err, bbbr)*/

	// Get bitcoin best bid fo USD
	rate, err := bp.GetBitcoinBestRate("USD")
	fmt.Println(err, rate)
}
