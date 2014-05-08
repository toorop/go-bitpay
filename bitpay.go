package bitpay

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	VERSION                    = 0.1
	API_ENDPOINT               = "https://bitpay.com/api/"
	DEFAULT_HTTPCLIENT_TIMEOUT = 30
)

// A bitpay represents a bitpay client wrapper
type bitpay struct {
	client *client
}

// New return a new Bitpay client
func New(apiKey string) *bitpay {
	return &bitpay{NewClient(apiKey)}
}

// GetBitcoinBestBidRates return Bitcoin Best Bid Rates (see bitpay doc)
func (b *bitpay) GetBitcoinBestBidRates() (bbbRates []bbbRate, err error) {
	r, err := b.client.do("GET", "rates", "")
	if err != nil {
		return
	}
	err = json.Unmarshal(r, &bbbRates)
	return
}

// GetBitcoinBestRate return bitcoin rate for currency currencyCode (helper)
func (b *bitpay) GetBitcoinBestRate(currencyCode string) (rate bbbRate, err error) {
	rates, err := b.GetBitcoinBestBidRates()
	if err != nil {
		return
	}
	err = errors.New(fmt.Sprintf("Currency %s not found", currencyCode))
	for _, rate = range rates {
		if rate.Code == currencyCode {
			err = nil
			break
		}
	}
	return
}
