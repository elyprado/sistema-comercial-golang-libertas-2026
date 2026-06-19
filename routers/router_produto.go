package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupProdutoRoutes(router *mux.Router) {
	router.HandleFunc("/produtos", controllers.GetProdutos).Methods("GET")
}
