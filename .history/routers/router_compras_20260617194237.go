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

	// Rota para retornar fornecedores e produtos para a criação de compras
	router.HandleFunc("/fornecedores", controllers.GetFornecedoresDisponiveis).Methods("GET")
	router.HandleFunc("/produtos", controllers.GetProdutosDisponiveis).Methods("GET")
}
