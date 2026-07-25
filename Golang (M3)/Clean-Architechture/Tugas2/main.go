package main

import (
	"fmt"
	"tugas/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/perpus", handler.LibHandler)
	// http.HandleFunc("/user", handler.UserHandler)

	fmt.Println("Server berjalan di port 8080")


	http.ListenAndServe(":8080", nil)
}