package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nopestack/cryptracker/service"
)

func CoinList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Printf("Requesting coin list\n")
	response := service.RequestCoins()

	rp := &Response{}

	rp.FromJSON(response)

	w.Header().Set("Content-Type", "application/json")
	//w.Write(response)
	encoded, _ := rp.ToBytes()
	w.Write(encoded)
}

func CoinDetail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}