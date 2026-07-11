package main

import (
	"perpus/library"
)

func main () {

	Buku := library.Book{
		Judul: "Bumi",
		Penulis : "Tere Liye",
	}

	library.Welcome()
	library.ShowBook(Buku);

}