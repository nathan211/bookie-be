package main

import (
	"bookie-be/db"
	"bookie-be/routes"
)

func main() {
	db.Init()
	r := routes.SetupRouter()
	r.Run(":8080")
}
