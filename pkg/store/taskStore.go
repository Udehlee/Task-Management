package store

import (
	"log"

	"github.com/Udehlee/Task-Management/pkg/models"
)

func (p PgConn) InsertTask(task models.Task) error {

	query := "INSERT INTO task(task_id,user_id, title,description) VALUES($1,$2,$3)"

	_, err := p.Conn.Exec(query, task.TaskID, task.UserID, task.Title, task.Description)
	if err != nil {
		log.Println("Error creating task:", err)
	}
	return nil

}

func (p PgConn) UpdateTask(task models.Task) error {

	query := "UPDATE tasks SET userId,description = $1, completed = $2 WHERE taskId = $3"

	_, err := p.Conn.Exec(query, task.UserID, task.Description, task.Completed, task.TaskID)
	if err != nil {
		log.Println("Error updating task:", err)
	}
	return err
}
