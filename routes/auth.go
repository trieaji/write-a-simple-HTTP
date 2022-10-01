package routes

import (
	"tchtest/handler"
	"tchtest/pkg/middleware"
	"tchtest/pkg/mysql"
	"tchtest/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handler.HandlerAuth(userRepository)

	router.HandleFunc("/register", h.Register).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}
