package store

import (
	"database/sql"
	"fmt"

	"github.com/Udehlee/Task-Management/pkg/models"
)

func (p PgConn) SaveUser(user models.User) error {

	query := "INSERT INTO users (firstname, lastname, email, pass_word) VALUES ($1, $2, $3, $4) RETURNING user_id, firstname, lastname, email"

	row := p.Conn.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)

	if err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return fmt.Errorf("error scanning row: %w", err)
	}

	return nil

}

func (p PgConn) UserByEmail(email string) (models.User, error) {

	query := `SELECT user_id, firstname, lastname, email, pass_word FROM users WHERE email = $1`
	var user models.User

	err := p.Conn.QueryRow(query, email).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No user found for email: %s\n", email) // Logging the error
			return models.User{}, fmt.Errorf("user not found")
		}
		fmt.Printf("Error querying user by email: %v\n", err) // Logging the error
		return models.User{}, fmt.Errorf("error querying user by email: %w", err)
	}

	return user, nil
}

func (p PgConn) GetAllUser() ([]models.User, error) {

	query := "SELECT user_id, firstname,lastname,email FROM users"

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

	query := "SELECT user_id, firstname,lastname,email FROM users WHERE user_id=$1"

	var user models.User
	err := p.Conn.QueryRow(query, id).Scan(&user.UserID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("error retrieving user")
		}
		return models.User{}, err
	}
	return user, nil
}
