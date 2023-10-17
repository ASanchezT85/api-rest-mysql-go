package routes

import (
	"github.com/ASanchezT85/api-rest-mysql-go/controllers"
	"github.com/gorilla/mux"
)

func SetUsersRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api/users").Subrouter()
	subRoute.HandleFunc("", controllers.Index).Methods("GET")
	subRoute.HandleFunc("", controllers.Store).Methods("POST")
	subRoute.HandleFunc("/{id}", controllers.Show).Methods("GET")
	subRoute.HandleFunc("/{id}", controllers.Destroy).Methods("DELETE")
}
