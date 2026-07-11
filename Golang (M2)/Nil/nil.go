package main

import "fmt"

func main() {

	var username string
	var logincoun int
	var status bool
	
	fmt.Println("username : ",username)
	fmt.Println("login count : ", logincoun)
	fmt.Println("status : ", status)

	var user []string
	fmt.Println(user)
	fmt.Println(user == nil)

	if user == nil {
		fmt.Println("belum ada data user")
	} else {
		fmt.Println("data user sudah ada")
	}

	user = append(user, "noir")
	fmt.Println(user)
	fmt.Println(user == nil)

}	