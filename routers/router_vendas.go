package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupVendaRoutes(router *mux.Router) {
	router.HandleFunc("/vendas", controllers.GetVendas).Methods("GET")
	router.HandleFunc("/vendas/{id}", controllers.GetVendaById).Methods("GET")
	router.HandleFunc("/vendas", controllers.CreateVenda).Methods("POST")
}
