package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ClientConfig struct {
	BaseURL    string
	APIKey     string
	APIVersion int
}

type CMCClient struct {
	baseURL      string
	apiVersion   int
	apiKey       string
	apiFormat    string
	apiHeaderKey string
	http.Client
}

func New(c ClientConfig) *CMCClient {
	// Craft an http client based on the config passed in
	client := &CMCClient{
		apiKey:       c.APIKey,
		apiVersion:   c.APIVersion,
		baseURL:      c.BaseURL,
		apiFormat:    "application/json",
		apiHeaderKey: "X-CMC_PRO_API_KEY",
	}

	return client
}

func (client CMCClient) request(method string, endpoint string) []byte {
	url := buildURL(client.baseURL, client.apiVersion, endpoint)

	req, err := http.NewRequest(
		method,
		url,
		nil,
	)

	if err != nil {
		log.Printf("[server] Error: %v\n", err)
	}

	req.Header.Set("Accepts", client.apiFormat)
	req.Header.Add(client.apiHeaderKey, client.apiKey)

	req.URL.RawQuery = encodeQuery(map[string]string{
		"start": "1",
	})

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("[server] Error: %v\n", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody
}

func (c CMCClient) GetCoins() []byte {
	// Read from DB, if data stale, request new
	endpoint := "cryptocurrency/listings/latest"
	return c.request("GET", endpoint)
}

func (c CMCClient) GetDetail() {
	// Read from DB, if data stale, request new
}
