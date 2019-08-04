package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func encodeQuery(values map[string]string) string {
	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	return q.Encode()
}

func RequestCoins() []byte {
	client := &http.Client{}

	method := "GET"
	baseURL := "https://pro-api.coinmarketcap.com/v1"
	route := "/cryptocurrency/listings/latest"

	req, err := http.NewRequest(method, fmt.Sprintf("%v%v", baseURL, route), nil)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "1")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("API_KEY"))

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	fmt.Println(resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error decoding response")
		os.Exit(1)
	}

	return respBody
}
