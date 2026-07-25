package main

import "fmt"

func login() {
	defer func() {
		if err := recover();
		err != nil {
			fmt.Println("Recovered : ", err)
		}
	}()

	fmt.Println("Program started")
	panic("Something went wrong")
	fmt.Println("Program finished")
}

func main() {
	login()
	fmt.Println("Program tetap jalan")
}