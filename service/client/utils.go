package client

import (
	"fmt"
	"net/url"
	"strconv"
)

func encodeQuery(values map[string]string) string {
	q := url.Values{}

	for key, value := range values {
		q.Add(key, value)
	}

	return q.Encode()
}

func buildURL(baseURL string, version int, route string) string {
	// mix endpoint, api version and base URL
	v := strconv.Itoa(version)

	return fmt.Sprintf("%v/v%v/%v", baseURL, v, route)
}
