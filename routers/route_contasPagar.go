package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupContasPagarRoutes(router *mux.Router) {
	router.HandleFunc("/contas-pagar", controllers.CreateContaPagar).Methods("POST")
	router.HandleFunc("/contas-pagar", controllers.GetContasPagar).Methods("GET")
	router.HandleFunc("/contas-pagar/{id}", controllers.GetContasPagarById).Methods("GET")
}
