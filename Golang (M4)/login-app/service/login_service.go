package service

import (
	"database/sql"
	"errors"
	"login/model"
	"login/repository"
)

func CreateUser(db *sql.DB, user model.User) error {

	if user.Username == "" {
		return errors.New("username tidak boleh kosong")
	}

	if user.Password == "" {
		return errors.New("password tidak boleh kosong")
	}

	return repository.InsertUser(db, user)

}

func GetAllUsers(db *sql.DB) ([]model.User, error) {

	return repository.GetAllUsers(db)

}

func GetUserById(db *sql.DB, id int) (model.User, error) {

	if id == 0 {
		return model.User{}, errors.New("id tidak boleh kosong")
	}

	return repository.GetUserById(db, id)

}

func UpdateUser (db *sql.DB, user model.User) error {
	
	if user.ID == 0{
		return errors.New("id tidak boleh kosong")
	}

	if user.Username == ""{
		return errors.New("username tidak boleh kosong")
	}

	if user.Password == ""{
		return errors.New("password tidak boleh kosong")
	}

	return repository.UpdateUser(db, user)

}

func DeleteUser(db *sql.DB, id int) error {

	if id == 0{
		return errors.New("id tidak boleh kosong")
	}

	return repository.DeleteUser(db, id)

}