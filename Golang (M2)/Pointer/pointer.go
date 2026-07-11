package main

import "fmt"

type User struct {
	Username string
	Password string
}

func gantiPassword(user *User) {
	user.Password = "12345"
}

func main() {

	user := User{
		Username: "admin",
		Password: "123",
	}
	
	fmt.Println("username :", user.Username)
	fmt.Println("password :", user.Password)

	gantiPassword(&user)

	fmt.Println("password setelah diubah :", user.Password)

	fmt.Println("username :", user.Username)
	fmt.Println("password :", user.Password)
	
	// username := "admin"

	// fmt.Println("username :", username)
	// fmt.Println("alamat memori variabel :", &username)

	// pointerUsername := &username

	// fmt.Println("alamat memori pointer :", pointerUsername)
	// fmt.Println("nilai pointer :", *pointerUsername)

}