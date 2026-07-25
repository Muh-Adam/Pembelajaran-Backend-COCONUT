package repository

type Data struct{
	NIM string
	Nama string
	Ipk string
}

func GetData() Data {
	return Data{
		NIM : "1234567",
		Nama : "orang",
		Ipk : "4.0",
	}
}