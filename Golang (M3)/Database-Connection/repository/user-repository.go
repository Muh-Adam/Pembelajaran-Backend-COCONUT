package repository

import (
	"app/database"
	"app/model"
)

func InsertUser(User model.User) error {

	db, err := database.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "INSERT INTO users (username, password) VALUES (?, ?)"

	_, err = db.Exec(query, User.Username, User.Password)

	if err != nil {
		return err
	}

	return nil

}

