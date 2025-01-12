package main

import (
	"backend/database"
	"backend/migrations"
	"backend/router"
	"fmt"
	"net/http"
)

func main() {
	database.Connect()
	migrations.Migrate()
	r := router.InitRouter()
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", r)
}
