// package main

// import "fmt"

// func cekUmur(umur int) {
// 	defer fmt.Println("Pemeriksaan selesai")
// 	defer func() {
// 		if umur < 17{
// 			if err := recover();
// 			err != nil {
// 			fmt.Println("Recovered : ", err)
// 			}
// 			defer panic("Umur belum mencukupi")
// 		} else {
// 			fmt.Println("Umur mencukupi")
// 		}
// 	}()
	
// }

// func main(){
// 	defer fmt.Println("Program selesai")
// 	cekUmur(13)	
// }