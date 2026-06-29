package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupCatalogRoutes(router *mux.Router) {
	router.HandleFunc("/clientescatalog", controllers.GetClientesCatalog).Methods("GET")
	router.HandleFunc("/vendedorescatalog", controllers.GetVendedoresCatalog).Methods("GET")
	router.HandleFunc("/produtoscatalog", controllers.GetProdutosCatalog).Methods("GET")
}
