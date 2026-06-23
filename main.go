package main

import (
	"log"
	"net/http"

	"go-crud-api/routers"

	"github.com/rs/cors"
    "go-crud-api/routers"
    "log"
    "net/http"
)

func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func main() {
	router := routers.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           43200,
	})

	handler := c.Handler(router)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8081", handler))
    router := routers.SetupRouter()
	routers.SetupClienteRoutes(router)
    log.Println("Starting server on :8080")
    
    log.Fatal(http.ListenAndServe(":8080", enableCORS(router)))
}
