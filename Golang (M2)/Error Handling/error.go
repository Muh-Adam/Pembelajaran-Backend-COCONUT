package main

import (
	"fmt"
	"errors"
)

func Login(username string, password string) error {
	if username != "admin" {
		return errors.New("Username salah")
	}

	if password != "12345"{
		return errors.New("Password salah")
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