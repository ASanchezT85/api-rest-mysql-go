package main

import (
	"log"
	"net/http"

	"github.com/ASanchezT85/api-rest-mysql-go/commons"
	"github.com/ASanchezT85/api-rest-mysql-go/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetUsersRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el prueto 9000")
	log.Println(server.ListenAndServe())
}
