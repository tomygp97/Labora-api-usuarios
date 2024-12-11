package main

import (
	"database/sql"
	"log"
	"net/http"
	"users-api/routes"
)

var db *sql.DB

func main() {
	routes.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/api-usuarios")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
