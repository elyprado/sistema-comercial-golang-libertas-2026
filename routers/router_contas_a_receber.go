package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupContaAReceberRoutes(router *mux.Router) {
	router.HandleFunc("/contas-a-receber", controllers.GetContasAReceber).Methods("GET")
	router.HandleFunc("/contas-a-receber", controllers.CreateContaAReceber).Methods("POST")
	router.HandleFunc("/contas-a-receber/{id}/baixar", controllers.BaixarContaAReceber).Methods("PUT")
}
