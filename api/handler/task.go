package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Udehlee/Task-Management/utils"
)

// AddUserTask adds task to user
func (h Handler) AddUserTask(w http.ResponseWriter, r *http.Request) {

	var task struct {
		UserID      int    `json:"username"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.AddTaskToUser(task.UserID, task.Title, task.Description, task.Completed)
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
		TaskID      int    `json:"task_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Completed   bool   `json:"completed"`
	}

	jd := json.NewDecoder(r.Body)
	if err := jd.Decode(&taskUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Service.UpdateUserTask(taskUpdate.TaskID, taskUpdate.Title, taskUpdate.Description, taskUpdate.Completed)
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
