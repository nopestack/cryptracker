package main

import (
	"os"
	"strconv"

	_ "github.com/nopestack/cryptracker/config"
	"github.com/nopestack/cryptracker/server"
)

func main() {

	_ = os.Getenv("API_KEY")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")

	c := server.ServerConfig{
		Address: host,
		Port:    port,
	}

	apiSrv := server.NewServer(c)
	apiSrv.Start()
}
