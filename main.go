package main

import (
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
	routers.SetupClienteRoutes(router)
    log.Println("Starting server on :8080")
    
    log.Fatal(http.ListenAndServe(":8080", enableCORS(router)))
}
