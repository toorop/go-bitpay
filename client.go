package bitpay

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// A cusher represents the HTTP client.
type client struct {
	apiKey string
	http.Client
}

// newClient return a new Bitpay client
func newClient(apiKey string) (c *client) {
	return &client{apiKey: apiKey}
}

// doTimeoutRequest process the request
func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Cryspy API")
	}
}

// do prepare and maje the request to bitpay API
func (c *client) do(method string, ressource string, payload string) (response []byte, err error) {
	connectTimer := time.NewTimer(DEFAULT_HTTPCLIENT_TIMEOUT * time.Second)

	query := fmt.Sprintf("%s%s", API_ENDPOINT, ressource)

	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.apiKey, "")

	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
