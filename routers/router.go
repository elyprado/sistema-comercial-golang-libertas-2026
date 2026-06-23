package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	// Configurar rotas de usuários em um arquivo separado
	SetupUsuarioRoutes(router)

	// Configurar rotas de contas a receber em um arquivo separado
	SetupContaAReceberRoutes(router)
	// Configaração rotas de compras em um arquivo separado
	SetupComprasRouters(router)

	// Configaração rotas de compraItem em um arquivo separado
	SetupCompraItemRouters(router)

	// Configurar rotas de fornecedor em um arquivo separado
	SetupFornecedorRoutes(router)

	// Configurar rotas de produto em um arquivo separado
	SetupProdutoRoutes(router)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return router
}
