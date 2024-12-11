package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"users-api/models"
)

var db *sql.DB
var users = []models.User{}

// var nextID = 1

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request, userId int) {
	for _, user := range users {
		if user.ID == userId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = int(id)
	user.CreatedAt = time.Now()
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Obtener usuarios")
}

func DeleteUser(w http.ResponseWriter, r *http.Request, id int) {
	log.Println("Obtener usuarios")
}
