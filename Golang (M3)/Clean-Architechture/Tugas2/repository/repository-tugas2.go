package repository

type Lib struct{
	Kode string
	Judul string
	Penulis string  
}

func GetData() Lib {
	return Lib{
		Kode : "BK001",
		Judul: "Pemrograman Go",
		Penulis: "Budi Santoso",
	}
}