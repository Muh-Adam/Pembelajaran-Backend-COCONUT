package main

import (
	"fmt"
	"app/handler"
	"net/http"
	"app/database"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Database tidak terkoneksi")
		fmt.Println(err)
		return
	}

	defer db.Close()

	http.HandleFunc("/login", handler.LoginHandler)
	// http.HandleFunc("/user", handler.UserHandler)

	fmt.Println("Server berjalan di port 8080")

	http.ListenAndServe(":8080", nil)

}