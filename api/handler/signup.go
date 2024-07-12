package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Udehlee/Task-Management/utils"
)

//Signup signs new user

func (h Handler) Signup(w http.ResponseWriter, r *http.Request) {

	var newUser struct {
		FirstName string `json:"firstname" validate:"required,alpha"`
		LastName  string `json:"lastname" validate:"required,alpha"`
		Email     string `json:"email" validate:"required,alpha"`
		Password  string `json:"password" validate:"required,alpha"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate user fields
	if err := utils.Validate(newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create user
	user, err := h.Service.CreateUser(newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "Registration unsuccessful", http.StatusBadRequest)
		return

	}

	AccessToken, err := utils.GenerateToken(user)
	if err != nil {
		http.Error(w, "Failed to generate token for sign up", http.StatusInternalServerError)
		return

	}

	//send this confirmation if signed up successfully
	SuccessfulRegMsg := map[string]interface{}{
		"message": "Signup successfully",
		"token":   AccessToken, //will be generated,
		"userinfo": map[string]string{
			"firstname": newUser.FirstName,
			"lastname":  newUser.LastName,
			"email":     newUser.Email,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessfulRegMsg)

}
