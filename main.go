package main

import (
	"fmt"
	"net/http"
	"tchtest/database"
	"tchtest/pkg/mysql"
	"tchtest/routes"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB here ...
	mysql.DatabaseInit()

	database.RunMigration()

	// On http (API)
	route := mux.NewRouter()

	routes.RouteInit(route.PathPrefix("/api/v1").Subrouter())

	route.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", route)
}
