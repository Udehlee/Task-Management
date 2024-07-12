package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Udehlee/Task-Management/pkg/models"
)

func (p PgConn) SaveUser(user models.User) error {

	userQuery := "INSERT INTO user (userId,firstname,lastname,email,password) VALUES($1,$2,$3,$4,$5) "

	_, err := p.Conn.Exec(userQuery, user.UserID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Fatal("error saving user to database")
	}
	return nil

}

func (p PgConn) UserByEmail(email string) (models.User, error) {

	CheckQuery := "SELECT userId,firstname,lastname,email,password FROM user WHERE email=$1 "
	var user models.User

	err := p.Conn.QueryRow(CheckQuery, email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found")
		}
		return models.User{}, err
	}
	return user, nil

}

func (p PgConn) GetAllUser() ([]models.User, error) {

	query := "SELECT userId, firstname,lastname,email FROM user"

	var users []models.User

	rows, err := p.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}

func (p PgConn) GetUserById(id int) (models.User, error) {

	query := "SELECT userId, firstname,lastname,email FROM user WHERE id=$1"

	var user models.User
	err := p.Conn.QueryRow(query, id).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("no user found")
		}
		return models.User{}, err
	}
	return user, nil
}
