package service

import (
	"tugas/repository"
)

func LibService(kode string, judul string, penulis string) string {

	Library := repository.GetData()

	if kode == ""{
		return "Kode tidak boleh kosong"
	}

	if kode != Library.Kode{
		return "Data buku tidak ditemukan"
	}

	return "Data ditemukan \n" + "Kode : " + Library.Kode + "\n" + "Judul : " + Library.Judul + "\n" + "Penulis : " + Library.Penulis 

}