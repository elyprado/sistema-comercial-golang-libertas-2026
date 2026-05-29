package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupCompraItemRouters(router *mux.Router) {
	router.HandleFunc("/compraItem", controllers.GetCompraItem).Methods("GET")
	router.HandleFunc("/compraItem/{id}", controllers.GetCompraItemById).Methods("GET")
	router.HandleFunc("/compraItem", controllers.CreateCompraItem).Methods("POST")
	router.HandleFunc("/compraItem/{id}", controllers.UpdateCompraItem).Methods("PUT")
	router.HandleFunc("/compraItem/{id}", controllers.DeleteCompraItem).Methods("DELETE")
}
