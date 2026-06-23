package routers

import (
	"go-crud-api/controllers"

	"github.com/gorilla/mux"
)

func SetupUsuarioRoutes(router *mux.Router) {
	router.HandleFunc("/usuarios", controllers.GetUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", controllers.GetUsuarioById).Methods("GET")
	router.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
	router.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
