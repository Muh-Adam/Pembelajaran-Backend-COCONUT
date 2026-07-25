package handler

import (
	"fmt"
	"net/http"
	"login/service"
)

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Login handler berhasil")
// }

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	result := service.LoginService(username, password)
	
	fmt.Fprintln(w, result)

}

// func UserHandler(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "User handler berhasil")

// }