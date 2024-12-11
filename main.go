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

	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/api-usuarios")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
