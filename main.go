package main

import (
	"log"
	"net/http"

<<<<<<< Updated upstream
	"github.com/rs/cors"
=======
	"go-crud-api/routers"
>>>>>>> Stashed changes
)

func main() {
	router := routers.SetupRouter()
<<<<<<< Updated upstream

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
	log.Fatal(http.ListenAndServe(":8080", handler))
=======
	log.Println("Starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
>>>>>>> Stashed changes
}
