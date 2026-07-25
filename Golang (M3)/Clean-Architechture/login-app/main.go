package main

import (
	"fmt"
	"login/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.LoginHandler)
	// http.HandleFunc("/user", handler.UserHandler)

	fmt.Println("Server berjalan di port 8080")

	http.ListenAndServe(":8080", nil)
}