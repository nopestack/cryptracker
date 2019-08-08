package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nopestack/cryptracker/server/router"
)

type ServerConfig struct {
	Port    int
	Address string
}

func (c ServerConfig) GetURL() string {
	return fmt.Sprintf("%v:%v", c.Address, c.Port)
}

type API struct {
	server http.Server
}

func NewServer(c ServerConfig) *API {
	mux := router.New()
	server := &API{
		server: http.Server{
			Addr:    c.GetURL(),
			Handler: mux,
		},
	}
	return server
}

func (a *API) Start() {
	log.Printf("[server] Initiating server at %v\n", a.server.Addr)

	log.Fatal(a.server.ListenAndServe())
}
