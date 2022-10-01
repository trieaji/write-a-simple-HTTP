package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(router *mux.Router) {
	Datas(router)
	Users(router)
	AuthRoutes(router)
}
