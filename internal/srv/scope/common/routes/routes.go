package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router, handler http.Handler, routes ...string) {
	for _, route := range routes {
		router.Handle(route, handler)
	}
}
