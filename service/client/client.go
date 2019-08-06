package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type ClientConfig struct {
	APIKey string
	APIVersion int
}

type CMCClient struct {
	baseURL string
	apiVersion int
	apiKey string
	apiFormat string
	apiHeaderKey string
	http.Client
}

func New(c ClientConfig) *CMCClient {
	// Craft an http client based on the config passed in
	client := &CMCClient{
		apiKey: c.APIKey,
		apiVersion: c.APIVersion,
		baseURL := "https://pro-api.coinmarketcap.com",
		apiFormat: "application/json",
		apiHeaderKey: "X-CMC_PRO_API_KEY",
		http.Client{},
	}

	return client

}

func (client CMCClient) request(method string, endpoint string) *http.Request {
	url := buildURL(client.baseURL, client.apiVersion, endpoint)

	req, err := http.NewRequest(
		method, 
		url,
		nil
	)

	if err != nil {
		log.Print(err)
		//os.Exit(1)
	}

	req.Header.Set("Accepts", client.apiFormat)
	req.Header.Add(client.apiHeaderKey, client.apiKey)

	req.URL.RawQuery = encodeQuery(map[string]string{
		"start": "1"
		//"limit": "1"
	})

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request to server")
		//os.Exit(1)
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