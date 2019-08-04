package client

import (
	"net/url"
)

func encodeQuery(values map[string]string) string {
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	return q.Encode()
}

func buildURL() string {
	// mix endpoint, api version and base URL
}