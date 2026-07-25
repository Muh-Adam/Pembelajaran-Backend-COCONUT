package handler

import (
	"fmt"
	"net/http"
	"tugas/service"
)

func LibHandler(w http.ResponseWriter, r *http.Request) {

	kode := r.URL.Query().Get("kode")
	judul := r.URL.Query().Get("judul")
	penulis := r.URL.Query().Get("penulis")

	result := service.LibService(kode, judul, penulis)
	
	fmt.Fprintln(w, result)

}