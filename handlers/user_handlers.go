package handlers

import (
	"net/http"
	"strconv"
	"users-api/controllers"
)

func HandleUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		controllers.GetUserById(w, r, id)
	case http.MethodPut:
		controllers.UpdateUser(w, r, id)
	// case http.MethodDelete:
	// 	controllers.DeleteUser(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
