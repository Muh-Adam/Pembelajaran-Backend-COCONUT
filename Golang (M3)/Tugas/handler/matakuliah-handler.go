package handler

import (
	"app/service"
	"fmt"
	"net/http"
)

func MataKuliahHandler(w http.ResponseWriter, r *http.Request) {
	kode := r.URL.Query().Get("kode")

	mk, err := service.CariMataKuliah(kode)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Kode: %s\nNama Mata Kuliah: %s\nSKS: %d\n", mk.Kode, mk.Nama, mk.SKS)
}
