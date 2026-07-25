package main

import "fmt"

func main() {
	defer fmt.Println("Program selesai")
	fmt.Println("Program dimulai")
	panic("Ada kesalahan!!")
	fmt.Println("Program selesai")
	defer fmt.Println("Akan selalu jalan!!")
}