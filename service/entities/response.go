package entities

import (
	"encoding/json"
	"fmt"
)

type ResponseStatus struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Data   []Coin         `json:"data"`
	Status ResponseStatus `json:"status"`
}

func (c *Response) FromJSON(s []byte) {

	err := json.Unmarshal(s, c)

	if err != nil {
		fmt.Printf("Error converting response to coin %v\n", err)
	}

	fmt.Printf("%v", c)
}

func (c Response) ToBytes() ([]byte, error) {

	encoded, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	return encoded, nil
}
