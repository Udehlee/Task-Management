package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Udehlee/Task-Management/utils"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {

	var userLogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&userLogin); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	user, err := h.Service.CheckUser(userLogin.Email, userLogin.Password)
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "user not found", http.StatusBadRequest)
		return
	}

	AccessToken, err := utils.GenerateToken(user)
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "failed to generate token", http.StatusBadRequest)
		return

	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Login successful",
		"data": map[string]any{
			"accessToken": AccessToken,
			"user": map[string]any{
				"userId":    user.UserID,
				"firstName": user.FirstName,
				"lastName":  user.LastName,
				"email":     user.Email,
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}
