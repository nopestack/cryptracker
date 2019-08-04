package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nopestack/cryptracker/service"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "API index")
}