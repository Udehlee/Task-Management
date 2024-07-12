package models

type Task struct {
	TaskID      int    `json:"task_id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"task_title"`
	Description string `json:"task_description"`
	Completed   bool   `json:"done"`
}
