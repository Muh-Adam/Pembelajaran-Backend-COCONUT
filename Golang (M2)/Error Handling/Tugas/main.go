package main

import (
	"fmt"
	"errors"
)

// func Login(pin int) error {

// 	if pin != 12345 {
// 		return errors.New("Akses Ditolak")
// 	}

// 	return nil

// }

// func main() {
// 	var input int

// 	for i := 1; i < 4; i++ {
// 	fmt.Println("Masukkan pin : ")
// 	fmt.Scan(&input)

// 		err := Login(input)
// 		if err != nil {
// 			fmt.Println(err)
// 			fmt.Println("Pin salah")
// 		} else {
// 			fmt.Println("Login berhasil")
// 			return
// 		}
		
// 	}
// }

func checkPin(pin string) error {
	for i := 0; i < 3;  i++ {
		fmt.Println("Masukkan PIN : ")
		fmt.Scan(&pin)
		if pin == "12345" {
			return nil
		}
		fmt.Println("Silahkan coba lagi")
	}
	return errors.New("PIN salah")
}

func main() {
	var pin string
	err := checkPin(pin)
	if err != nil {
		fmt.Println("Akses Ditolak")
		fmt.Println("Penyebab : ", err)
	} else {
		fmt.Println("Login berhasil")
	}
}