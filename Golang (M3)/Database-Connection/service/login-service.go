package service

import (
	"app/repository"
	"errors"
	"app/model"
)

func LoginService(username string, password string) string {

	user := repository.GetUser()

	if username == "" {
		return "Username tidak boleh kosong"
	}

	if password == "" {
		return "Password tidak boleh kosong"
	}

	if username != user.Username{
		return "Username salah"
	}

	if password != user.Password{
		return "Password salah"
	}

	return "Login berhasil"

}
func CreateUserService(User model.User) error {

	user := model.User{}

	if username == "" {
		return errors.New("Username tidak boleh kosong")
	}

	if password == "" {
		return errors.New("Password tidak boleh kosong")
	}

	repository.InsertUser(user)

	return nil
}