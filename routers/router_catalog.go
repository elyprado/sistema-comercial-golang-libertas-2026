package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupCatalogRoutes(router *mux.Router) {
	router.HandleFunc("/clientes", controllers.GetClientes).Methods("GET")
	router.HandleFunc("/vendedores", controllers.GetVendedores).Methods("GET")
	router.HandleFunc("/produtos", controllers.GetProdutos).Methods("GET")
}
