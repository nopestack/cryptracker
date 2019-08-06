package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nopestack/cryptracker/service/client"
	"github.com/nopestack/cryptracker/service/entities"
)

var cmcConfig = client.ClientConfig{
	APIKey:     os.Getenv("API_KEY"),
	APIVersion: 1,
}

var cmcClient = client.New(cmcConfig)

func CoinList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("Requesting coin list\n")
	response := cmcClient.GetCoins()

	rp := &entities.Response{}

	rp.FromJSON(response)

	w.Header().Set("Content-Type", "application/json")
	//w.Write(response)
	encoded, _ := rp.ToBytes()
	w.Write(encoded)
}

func CoinDetail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}
