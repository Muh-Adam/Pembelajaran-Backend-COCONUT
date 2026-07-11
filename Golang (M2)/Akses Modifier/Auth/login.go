package auth

import "fmt"

type User struct {
	Username string
	Password string
}

func Login(user User, username string, password string) {

	if username == user.Username && password == user.Password {
		fmt.Println("Login Success")
	} else {
		fmt.Println("Login Failed")
	}
	
}

func Info() {
	fmt.Println("Package Auth berhasil digunakan")
}

func checkPassword() {
	
}