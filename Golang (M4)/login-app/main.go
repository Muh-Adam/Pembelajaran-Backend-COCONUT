package main

import (
	"fmt"
	"login/database"
	"login/handler"
)

func main() {

	db, err := database.ConnectDB()

	if err != nil {
		fmt.Println("Koneksi database gagal :", err)
		return
	}

	defer db.Close()

	for {

		var menu int

		fmt.Println()
		fmt.Println("========================")
		fmt.Println("          MENU")
		fmt.Println("========================")
		fmt.Println("1. Tambah Data User")
		fmt.Println("2. Tampilkan Semua User")
		fmt.Println("3. Cari User Berdasarkan ID")
		fmt.Println("4. Update Data User")
		fmt.Println("5. Delete Data User")
		fmt.Println("6. Keluar")

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&menu)

		switch menu {

		case 1:

			handler.CreateUserHandler(db)

		case 2:
			handler.GetAllUsersHandler(db)

		case 3:

			handler.GetUserIdHandler(db)

		case 4:

			handler.UpdateUserHandler(db)

		case 5:

			handler.DeleteUserHandler(db)

		case 6:

			fmt.Println("Program selesai.")
			return

		default:

			fmt.Println("Menu tidak tersedia.")

		}

	}

}