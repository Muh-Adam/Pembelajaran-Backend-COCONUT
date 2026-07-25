package service

import (
	"tugas/repository"
)

func DataService(nim string, nama string, ipk string) string {

	Datas := repository.GetData()

	if nim == ""{
		return "Nim tidak boleh kosong"
	}

	if nim != Datas.NIM{
		return "Data tidak ditemukan"
	}

	if len(nim) < 7{
		return "format NIM salah"
	}

	

	return "Data ditemukan \n" + "NIM : " + Datas.NIM + "\n" + "Nama : " + Datas.Nama + "\n" + "IPK : " + Datas.Ipk 

}