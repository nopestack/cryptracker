package entities

type Currency struct {
	Price       float64 `json:"price"`
	LastUpdated string  `json:"last_updated"`
}

type Quote map[string]Currency

type Coin struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MaxSupply   int    `json:"max_supply"`
	LastUpdated string `json:"last_updated"`
	Quote       Quote  `json:"quote"`
}