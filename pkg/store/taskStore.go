package store

import (
	"log"

	"github.com/Udehlee/Task-Management/pkg/models"
)

func (p PgConn) InsertTask(task models.Task) error {
	query := "INSERT INTO tasks(user_id, title, descrip_tion, completed) VALUES($1, $2, $3, $4)"

	_, err := p.Conn.Exec(query, task.UserID, task.Title, task.Description, task.Completed)
	if err != nil {
		log.Println("Error creating task:", err)
		return err
	}
	log.Println("Task created successfully:", task.Title)
	return nil
}

// func (p PgConn) UpdateTask(task models.Task) error {

// 	query := "UPDATE tasks SET task_id,description = $1, completed = $2 WHERE taskId = $3"

// 	_, err := p.Conn.Exec(query, task.TaskID, task.Description, task.Completed, task.TaskID)
// 	if err != nil {
// 		log.Println("Error updating task:", err)
// 	}
// 	return err
// }

func (p PgConn) UpdateTask(task models.Task) error {

	query := " UPDATE tasks  SET title = $2, descrip_tion = $3, completed = $4 WHERE task_id = $1"

	_, err := p.Conn.Exec(query, task.TaskID, task.Title, task.Description, task.Completed)
	if err != nil {
		log.Println("Error updating task:", err)
		return err
	}
	return nil
}
