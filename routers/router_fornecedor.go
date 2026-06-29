package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupFornecedorRoutes(router *mux.Router) {
	router.HandleFunc("/fornecedores", controllers.GetFornecedores).Methods("GET")
	router.HandleFunc("/fornecedores/{id}", controllers.GetFornecedorById).Methods("GET")
	router.HandleFunc("/fornecedores", controllers.CreateFornecedor).Methods("POST")
	router.HandleFunc("/fornecedores/{id}", controllers.UpdateFornecedor).Methods("PUT")
	router.HandleFunc("/fornecedores/{id}", controllers.DeleteFornecedor).Methods("DELETE")
}
