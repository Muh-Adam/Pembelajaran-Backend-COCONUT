package handler

import (
	"fmt"
	"net/http"
	"tugas/service"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {

	nim := r.URL.Query().Get("nim")
	nama := r.URL.Query().Get("nama")
	ipk := r.URL.Query().Get("ipk")

	result := service.DataService(nim, nama, ipk)
	
	fmt.Fprintln(w, result)

}