package routes

import (
	"net/http"
	"users-api/controllers"
)

func SetupRouter() {
	http.HandleFunc("GET /users", controllers.GetAllUsers)
	http.HandleFunc("GET /users/{id}", controllers.GetAllUsers)
	http.HandleFunc("POST /users", controllers.CreateUser)
	// http.HandleFunc("PUT /users/{id}", controllers.UpdateUser)
	// http.HandleFunc("PUT /users/{id}", controllers.DeleteUser)
}
