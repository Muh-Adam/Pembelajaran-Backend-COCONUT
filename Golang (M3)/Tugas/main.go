package main

import (
	"app/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/matakuliah", handler.MataKuliahHandler)

	fmt.Println("Server berjalan di port 8080")

	http.ListenAndServe(":8080", nil)
}
