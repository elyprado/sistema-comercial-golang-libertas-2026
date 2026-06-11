package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupClienteRoutes(router *mux.Router) {
	router.HandleFunc("/clientes", controllers.GetClientes).Methods("GET")
	router.HandleFunc("/clientes/{id}", controllers.GetClienteById).Methods("GET")
	router.HandleFunc("/clientes", controllers.CreateCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", controllers.UpdateCliente).Methods("PUT")
	router.HandleFunc("/clientes/{id}", controllers.DeleteCliente).Methods("DELETE")
}
