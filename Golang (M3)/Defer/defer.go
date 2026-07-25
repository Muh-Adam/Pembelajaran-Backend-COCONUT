package main

import "fmt"

// func login() {
// 	fmt.Println("Memulai proses login")

// 	defer fmt.Println("Proses login selesai!")

// 	fmt.Println("Memeriksa username dan password")
// }

// func main() {
// 	login()
// }


func login(username string) {

	defer fmt.Println("Proses login selesai!")
	defer fmt.Println("memeriksa username dan password")
	defer fmt.Println("Memulai proses login")
	

	if username == ""{
		fmt.Println("Username kosong")
		return
	}

	fmt.Println("Memulai proses login")

}

func main() {
	defer login("")
	defer fmt.Println("Main selesai")
}

