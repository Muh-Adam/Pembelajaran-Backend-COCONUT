package main

import "fmt"

func tarikSaldo(saldo, jumlah int){
	defer func(){
			if err := recover();
			err != nil {
				fmt.Println("Transaksi ditolak : ", err)
			}
		}()
	defer fmt.Println("Rekening di cek")
	if jumlah > saldo{
		panic("saldo tidak mencukupi")
	}
}

func main(){
	tarikSaldo(50000, 100000)
	fmt.Println("Terima kasih karena sudah menggunakan layanan kami")
}