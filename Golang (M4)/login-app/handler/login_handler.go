package handler

import (
	"database/sql"
	"fmt"
	"login/model"
	"login/service"
)

func CreateUserHandler(db *sql.DB) {

	var user model.User

	fmt.Println()
	fmt.Println("===== TAMBAH DATA USER =====")

	fmt.Print("Username : ")
	fmt.Scan(&user.Username)

	fmt.Print("Password : ")
	fmt.Scan(&user.Password)

	err := service.CreateUser(db, user)

	if err != nil {

		fmt.Println("Error :", err)

		return

	}

	fmt.Println("Data berhasil ditambahkan")

}

func GetAllUsersHandler(db *sql.DB){
	
	users, err := service.GetAllUsers(db)
	
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println()
	fmt.Println("===== DAFTAR USER =====")

	for _, user := range users {
		fmt.Printf("ID : %d\n", user.ID)
		fmt.Printf("Username : %s\n", user.Username)
		fmt.Printf("Password : %s\n", user.Password)
		fmt.Println()
	}
}

func GetUserIdHandler(db *sql.DB) {

	var id int

	fmt.Println()
	fmt.Println("===== CARI USER BERDASARKAN ID =====")

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	user, err := service.GetUserById(db, id)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println()
	fmt.Println("===== DATA USER =====")
	fmt.Printf("ID : %d\n", user.ID)
	fmt.Printf("Username : %s\n", user.Username)
	fmt.Printf("Password : %s\n", user.Password)
	fmt.Println()

}

func UpdateUserHandler (db *sql.DB) {
	
	var user model.User

	fmt.Println()
	fmt.Println("===== UPDATE DATA USER =====")

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&user.ID)

	fmt.Print("Username : ")
	fmt.Scan(&user.Username)

	fmt.Print("Password : ")
	fmt.Scan(&user.Password)

	err := service.UpdateUser(db, user)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Data berhasil diupdate")

}

func DeleteUserHandler(db *sql.DB) {

	var id int

	fmt.Println()
	fmt.Println("===== DELETE DATA USER =====")

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	err := service.DeleteUser(db, id)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("Data berhasil dihapus")

}