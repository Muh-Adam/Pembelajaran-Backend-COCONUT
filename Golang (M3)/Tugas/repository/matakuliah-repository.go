package repository

import "app/model"

func GetMataKuliahs() []model.MataKuliah {
	return []model.MataKuliah{
		{Kode: "IF201", Nama: "Struktur Data", SKS: 3},
		{Kode: "IF202", Nama: "Basis Data", SKS: 3},
		{Kode: "IF203", Nama: "Pemrograman Web", SKS: 3},
		{Kode: "IF204", Nama: "Jaringan Komputer", SKS: 3},
		{Kode: "IF205", Nama: "Sistem Operasi", SKS: 3},
	}
}
