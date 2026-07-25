// package main

// import "fmt"

// func validasiPassword(pass string){

// 	defer func(){
// 			if err := recover();
// 			err != nil {
// 				fmt.Println("Validasi gagal : ", err)
// 			}
// 		}()
// 	defer fmt.Println("Log validasi tersimpan")
// 	if len(pass) < 6{
// 		panic("Password terlalu pendek")
// 	}


// }

// func main(){
// 	defer fmt.Println("Silahkan coba lagi")
// 	validasiPassword("tes")
	
// }