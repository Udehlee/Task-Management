package service

import (
	"fmt"
	"log"

	"github.com/Udehlee/Task-Management/pkg/models"
	"github.com/Udehlee/Task-Management/utils"
)

func (s Service) CreateUser(firstname, lastname, email, password string) (models.User, error) {

	hashedpwd, err := utils.HashPassword(password)
	if err != nil {
		log.Fatal("failed to hash password")
	}

	U := models.User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  hashedpwd,
	}

	if err := s.Store.SaveUser(U); err != nil {
		return models.User{}, fmt.Errorf("error saving user: %v", err)
	}

	// Return the created user object
	return U, nil

}

func (s Service) CheckUser(email, password string) (models.User, error) {

	user, err := s.Store.UserByEmail(email)
	if err != nil {
		return models.User{}, fmt.Errorf("user not found")

	}

	err = utils.ComparePasswordHash(user.Password, password)
	if err != nil {
		return models.User{}, fmt.Errorf("wrong password")
	}

	return models.User{}, nil

}

func (s Service) GetAllUser() ([]models.User, error) {

	users, err := s.Store.GetAllUser()
	if err != nil {
		return nil, fmt.Errorf("error get users")
	}
	return users, nil

}

func (s Service) GetUserById(id int) (models.User, error) {
	user, err := s.Store.GetUserById(id)
	if err != nil {
		return models.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}
