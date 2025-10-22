package main

import (
	"log"

	"github.com/reynaldodomenico/react-and-go-notes/go-backend/db"
	"github.com/reynaldodomenico/react-and-go-notes/go-backend/router"
)

func main() {
	db.Connect()

	r := router.SetupRouter()

	log.Println("🚀 Server running on :8080")
	r.Run(":8080")
}
