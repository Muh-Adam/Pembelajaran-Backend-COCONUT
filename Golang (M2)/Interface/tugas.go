package main

import "fmt"

type student interface {
	validasi(nim string, jurusan string)
}

type mahasiswa struct {
	nim string
	jurusan string
}

func (m mahasiswa) validasi (nim string, jurusan string){
	if m.nim == nim && m.jurusan == jurusan {
		fmt.Println ("Data mahasiswa valid")
	} else {
		fmt.Println ("Data mahasiswa tidak valid")
	}
}

func main (){
	
	mhs := mahasiswa{
		nim: "123",
		jurusan: "Teknik Informatika",
	}
	mhs.validasi("123", "Teknik Informatika")

}