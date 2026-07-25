package main

import (
	"fmt"
	"login/database"
	"login/handler"
	"net/http"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Koneksi database gagal :", err)
		return
	}
	defer db.Close()

	// Routing REST API
	http.HandleFunc("/users", handler.UserHandler(db))
	http.HandleFunc("/users/detail", handler.GetUserIdHandler(db))

	fmt.Println("Server REST API berjalan di http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal menjalankan server REST API:", err)
	}
}