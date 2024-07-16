package service

import "github.com/Udehlee/Task-Management/pkg/models"

//AddTaskToUser adds task specific user
func (s Service) AddTaskToUser(userID int, title, description string, completed bool) error {

	t := models.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	if err := s.Store.InsertTask(t); err != nil {
		return err
	}
	return nil
}

//AddTaskToUser updates  specific user task
func (s Service) UpdateUserTask(taskId int, title, description string, completed bool) error {

	Update := models.Task{
		TaskID:      taskId,
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	if err := s.Store.UpdateTask(Update); err != nil {
		return err
	}
	return nil
}
