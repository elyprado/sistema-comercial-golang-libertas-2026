package main

import (
	"go-crud-api/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.SetupRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
