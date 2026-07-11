package main

import "project/Auth"

func main () {

	auth.Info()

	user := auth.User{
		Username: "admin",
		Password: "123",
	}

	auth.Login(user, "admin", "123")
}


