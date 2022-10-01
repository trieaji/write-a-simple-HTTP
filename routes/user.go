package routes

import (
	"tchtest/handler"
	"tchtest/pkg/mysql"
	"tchtest/repositories"

	"github.com/gorilla/mux"
)

func Users(router *mux.Router) {
	PakekUserRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerUser(PakekUserRepository)

	router.HandleFunc("/users", h.FindUsers).Methods("GET")
	router.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	router.HandleFunc("/user", h.CreateUser).Methods("POST")
}
