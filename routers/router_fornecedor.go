package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupFornecedorRoutes(router *mux.Router) {
	router.HandleFunc("/fornecedores", controllers.GetFornecedores).Methods("GET")
}
