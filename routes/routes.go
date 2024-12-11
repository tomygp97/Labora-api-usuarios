package routes

import (
	"net/http"
	"users-api/controllers"
	"users-api/handlers"
)

func SetupRouter() {
	http.HandleFunc("GET /users", controllers.GetAllUsers)
	http.HandleFunc("GET /users/{id}", handlers.HandleUserById)
	http.HandleFunc("POST /users", controllers.CreateUser)
	http.HandleFunc("PUT /users/{id}", handlers.HandleUserById)
	http.HandleFunc("DELETE /users/{id}", controllers.DeleteUser)
}
