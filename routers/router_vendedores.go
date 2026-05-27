package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupVendedorRoutes(router *mux.Router) {
	router.HandleFunc("/vendedores", controllers.GetVendedores).Methods("GET")
	router.HandleFunc("/vendedores/{id}", controllers.GetVendedorById).Methods("GET")
	router.HandleFunc("/vendedores", controllers.CreateVendedor).Methods("POST")
	router.HandleFunc("/vendedores/{id}", controllers.UpdateVendedor).Methods("PUT")
	router.HandleFunc("/vendedores/{id}", controllers.DeleteVendedor).Methods("DELETE")
}
