package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupUsuarioRoutes(router *mux.Router) {
	router.HandleFunc("/usuarios", controllers.GetFornecedores).Methods("GET")
	router.HandleFunc("/usuarios/{id}", controllers.GetFornecedorById).Methods("GET")
	//-----------------------------------------------------------------------------------
	router.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
	router.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
}
