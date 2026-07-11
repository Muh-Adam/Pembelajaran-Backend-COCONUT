package main

import "fmt"

type LoginError struct {
	Message string
}

func (e LoginError) Error() string {
	return e.Message
}

func Login(username string, password string) error {
	if username == "" {
		return LoginError{
			Message: "Username tidak boleh kosong",
		}
	}

	if password == "" {
		return LoginError{
			Message: "Password tidak boleh kosong",
		}
	}

	if username != "admin" {
		return LoginError{
			Message: "username salah",
		}
	}

	if password != "12345" {
		return LoginError{
			Message: "Password salah",
		}
	}

	return nil
}

func main() {
	err := Login("admin", "12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Login berhasil")
}
	