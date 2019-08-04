package router

import (
	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	// Route definitions go here
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/coins/", CoinList)
	router.GET("/coins/:id", Index)
	return router
}
