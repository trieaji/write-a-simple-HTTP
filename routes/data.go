package routes

import (
	"tchtest/handler"
	"tchtest/pkg/middleware"
	"tchtest/pkg/mysql"
	"tchtest/repositories"

	"github.com/gorilla/mux"
)

func Datas(router *mux.Router) {
	UseDataRepository := repositories.RepositoryDataImage(mysql.DB)
	h := handler.HandlerDataImage(UseDataRepository)

	router.HandleFunc("/datas", h.FindDatas).Methods("GET")
	router.HandleFunc("/data/{id}", h.GetData).Methods("GET")
	router.HandleFunc("/data", middleware.Auth(middleware.UploadFile(h.CreateData))).Methods("POST")
}
