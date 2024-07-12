package service

import "github.com/Udehlee/Task-Management/pkg/models"

// func (s Service) AssignTask()
func (s Service) AddTaskToUser(userID int, title, description string) error {

	t := models.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	if err := s.Store.InsertTask(t); err != nil {
		return err
	}
	return nil
}

func (s Service) UpdateUserTask(userId, taskId int, title, description string, completed bool) error {
	// return s.Store.UpdateTask(taskId, description, completed)

	Update := models.Task{
		UserID:      userId,
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
