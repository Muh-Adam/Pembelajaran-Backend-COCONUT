package service

import "login/repository"

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