package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type CMCClient struct {
	apiKey string
	baseURL string
	apiVersion int
	client http.Client
}

func New(c ClientConfig) *CMCClient {
	// Craft an http client based on the config passed in
}

func (c CMCClient) GetCoins() {
	// Read from DB, if data stale, request new
}

func (c CMCClient) GetDetail() {
	// Read from DB, if data stale, request new
}