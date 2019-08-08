package router

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nopestack/cryptracker/service/client"
	"github.com/nopestack/cryptracker/service/entities"
)

var cmcConfig = client.ClientConfig{
	BaseURL:    os.Getenv("API_BASE_URL"),
	APIKey:     os.Getenv("API_KEY"),
	APIVersion: 1,
}

var cmcClient = client.New(cmcConfig)

func CoinList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("[server] %v %v %v\n", r.Method, r.URL, r.RemoteAddr)
	log.Printf("[server] Requesting coin list\n")
	response := cmcClient.GetCoins()
	log.Printf("[server] Coin list retrieved\n")

	rp := &entities.Response{}

	// NOTE: Marshalling/Unmarshalling could probably be handled better
	rp.FromJSON(response)

	w.Header().Set("Content-Type", "application/json")
	encoded, _ := rp.ToBytes()
	w.Write(encoded)
}

func CoinDetail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
