package service

import (
	"fmt"
	"log"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/Udehlee/Task-Management/utils"
)

// CreateUser creates User details
// saves it to the database
func (s Service) CreateUser(firstname, lastname, email, password string) (models.User, error) {

	hashedpwd, err := utils.HashPassword(password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := models.User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  hashedpwd,
	}

	if err := s.Store.SaveUser(user); err != nil {
		log.Printf("Error saving user: %v", err)
		return models.User{}, fmt.Errorf("error saving user to database: %w", err)
	}

	// Return the created user object
	return user, nil

}

// CheckUser checks if users is present
func (s *Service) CheckUser(email, password string) (models.User, error) {
	user, err := s.Store.UserByEmail(email)
	if err != nil {
		fmt.Printf("Error in UserByEmail: %v\n", err) // Logging the error
		return models.User{}, fmt.Errorf("user not found")
	}

	err = utils.ComparePasswordHash(user.Password, password)
	if err != nil {
		fmt.Printf("Error in ComparePasswordHash: %v\n", err) // Logging the error
		return models.User{}, fmt.Errorf("wrong password")
	}

	return user, nil
}

// GetAllUser retrieves all users
func (s Service) GetAllUser() ([]models.User, error) {

	users, err := s.Store.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("error get users")
	}
	return users, nil

}

// GetUserById retrieves a single user based on id
func (s *Service) GetUserById(id int) (models.User, error) {
	user, err := s.Store.GetUserById(id)
	if err != nil {
		return models.User{}, fmt.Errorf("user not found: %v", err)
	}
	return user, nil
}
