package main

import (
	"log"
	"net/http"

	"go-crud-api/routers"

	"github.com/rs/cors"
)

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
}
