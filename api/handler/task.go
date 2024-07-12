package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Udehlee/Task-Management/utils"
)

func (h Handler) AddUserTask(w http.ResponseWriter, r *http.Request) {

	var task struct {
		UserId      int    `json:"user_Id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Done        bool   `json:"bool"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.AddTaskToUser(task.UserId, task.Title, task.Description)
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "failed to add task", http.StatusInternalServerError)
		return
	}

	//successfully added task to user
	response := map[string]string{
		"message": "successfully added task to user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

// update when user complete task
func (h Handler) UpdateUserTask(w http.ResponseWriter, r *http.Request) {

	var taskUpdate struct {
		UserID      int    `json:"userId"`
		TaskID      int    `json:"taskid"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&taskUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add task
	err := h.Service.UpdateUserTask(taskUpdate.UserID, taskUpdate.TaskID, taskUpdate.Title, taskUpdate.Description, taskUpdate.Completed)
	if err != nil {
		utils.UnsucessfulRequest(w, "Bad Request", "failed to update task", http.StatusInternalServerError)
		return
	}

	//successfully added task to user
	response := map[string]string{
		"message": "successfully updated task to user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
