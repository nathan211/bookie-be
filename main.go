package main

import (
	"bookie-be/db"
	"bookie-be/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()
	r := routes.SetupRouter()
	r.Run(":8080")

	log.Fatal("Server is running...")
}
