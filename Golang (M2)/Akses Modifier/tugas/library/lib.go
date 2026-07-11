package library

import "fmt"

type Book struct {
	Judul   string
	Penulis string
}


func ShowBook(Buku Book) {
	fmt.Println("Judul Buku : ", Buku.Judul)
	fmt.Println("Penuliss Buku : ", Buku.Penulis)
}

func Welcome() {
	fmt.Println("Selamat datang di sistem perpustakaan")
}