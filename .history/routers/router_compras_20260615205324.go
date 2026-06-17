package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupComprasRouters(router *mux.Router) {
	router.HandleFunc("/compras", controllers.GetCompras).Methods("GET")
	router.HandleFunc("/compras/{id}", controllers.GetComprasById).Methods("GET")
	router.HandleFunc("/compras", controllers.CreateCompras).Methods("POST")
	router.HandleFunc("/compras/{id}", controllers.UpdateCompras).Methods("PUT")
	router.HandleFunc("/compras/{id}", controllers.DeleteCompras).Methods("DELETE")
}
