package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Udehlee/Task-Management/utils"
)

// GetAlluser returns all users in a json format
func (h Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {

	users, err := h.Service.GetAllUser()
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "unable to retrieve all users", http.StatusInternalServerError)
		return

	}

	UserJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(UserJSON)

}

func (h Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	userIDStr := parts[len(parts)-1]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	UserId, err := h.Service.GetUserById(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	UserIdJSON, err := json.Marshal(UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(UserIdJSON)

}
