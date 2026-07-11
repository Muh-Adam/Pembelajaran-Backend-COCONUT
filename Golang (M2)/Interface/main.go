// // Implementasi Interface


// package main

// import "fmt"

// type Login interface {
// 	Login(username string, password string)
// }

// type UserLogin struct {
// 	Username string
// 	Password string
// }

// type EmailLogin struct {
// 	Email string
// 	Password string
// }

// func (u UserLogin) Login(username string, password string) {
// 	if u.Username == username && u.Password == password {
// 		fmt.Println("Login Success")
// 	} else {
// 		fmt.Println("Login Failed")
// 	}
// }

// func (e EmailLogin) Login(username string, password string) {
// 	if e.Email == username && e.Password == password {
// 		fmt.Println("Login Success")
// 	} else {
// 		fmt.Println("Login Failed")
// 	}
// }

// func main() {

// 	var sistem Login
	
// 	user := UserLogin{
// 		Username: "admin",
// 		Password: "12345",
// 	}
// 	user.Login("admin", "12345")


// 	email := EmailLogin{
// 		Email: "[EMAIL_ADDRESS]",
// 		Password: "12345",
// 	}
	
// 	sistem = user
// 	sistem.Login("admin", "12345")

	
// 	sistem = email
// 	sistem.Login("[EMAIL_ADDRESS]", "12345")

	
// }
