package repository

import (
	"database/sql"
	"login/model"
)

func InsertUser(db *sql.DB, user model.User) error {

	query := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err := db.Exec(
		query,
		user.Username,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil

}

func GetAllUsers(db  *sql.DB) ([]model.User, error) {
	query := "SELECT id, username, password FROM users"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []model.User

	for rows.Next(){
		var user model.User 

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil

}

func GetUserById(db *sql.DB, id int) (model.User, error) {
	
	query := "SELECT id, username, password FROM users WHERE id = ?"

	var user model.User

	err := db.QueryRow(query, id).Scan(
		&user.ID, 
		&user.Username, 
		&user.Password,
	)

	if err != nil {
		return model.User{}, err
	}

	return user, nil

}

func UpdateUser (db *sql.DB, user model.User) error {
	
	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"

	_, err := db.Exec(
		query,
		user.Username,
		user.Password,
		user.ID,
	)

	if err != nil {
		return err
	}

	return nil

}

func DeleteUser(db *sql.DB, id int) error {
	
	query := "DELETE FROM users WHERE id = ?"

	_, err := db.Exec(
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
	
}