package main

import (
	"log"
	"net/http"

	"github.com/reynaldodomenico/react-and-go-notes/go-backend/db"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/router"
)

func main() {
	db.Connect()
	router.RegisterRoutes()

	log.Println("🚀 Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
