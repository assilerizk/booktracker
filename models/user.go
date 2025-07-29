package models

import (
	"booktracker/db"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

// CreateUser inserts a new user and returns the new user ID.
func CreateUser(username, password string) (int, error) {
	var id int
	err := db.DB.QueryRow(
		`INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`,
		username, password,
	).Scan(&id)
	return id, err
}

// GetUserByUsername fetches a user by username or returns an error.
func GetUserByUsername(username string) (User, error) {
	var user User
	err := db.DB.QueryRow(
		`SELECT id, username, password FROM users WHERE username = $1`,
		username,
	).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return User{}, errors.New("user not found")
	}
	return user, nil
}