package handler

import (
	"database/sql"
	"encoding/json"
	"login/model"
	"login/service"
	"net/http"
	"strconv"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateUserHandler(db, w, r)
			return
		}

		if r.Method == http.MethodGet {
			GetAllUsersHandler(db, w, r)
			return
		}

		if r.Method == http.MethodPut {
			UpdateUserHandler(db, w, r)
			return
		}

		if r.Method == http.MethodDelete {
			DeleteUserHandler(db, w, r)
			return
		}

		SendResponse(w, http.StatusMethodNotAllowed, Response{
			Message: "Method tidak diizinkan",
		})
	}
}

func CreateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{
			Message: "request invalid",
		})
		return
	}

	err = service.CreateUser(db, user)
	if err != nil {
		SendResponse(w, http.StatusCreated, Response{
			Message: err.Error(),
		})
		return
	}

	SendResponse(w, http.StatusOK, Response{
		Message: "Data berhasil ditambahkan",
	})
}

func GetAllUsersHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	users, err := service.GetAllUsers(db)

	if err != nil {
		SendResponse(w, http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
		return
	}

	SendResponse(w, http.StatusOK, Response{
		Message: "Data berhasil diambil",
		Data:    users,
	})
}

func GetUserIdHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{
				Message: "id harus berupa angka",
			})
			return
		}

		user, err := service.GetUserById(db, id)
		if err != nil {
			SendResponse(w, http.StatusNotFound, Response{
				Message: "user tidak ditemukan",
			})
			return
		}

		SendResponse(w, http.StatusOK, Response{
			Message: "Data berhasil diambil",
			Data:    user,
		})
	}
}

func UpdateUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{
			Message: "request invalid",
		})
		return
	}

	err = service.UpdateUser(db, user)

	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}

	SendResponse(w, http.StatusOK, Response{
		Message: "Data berhasil diupdate",
	})

}

func DeleteUserHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{
			Message: "id harus berupa angka",
		})
		return
	}

	err = service.DeleteUser(db, id)

	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{
			Message: err.Error(),
		})
		return
	}

	SendResponse(w, http.StatusOK, Response{
		Message: "Data berhasil dihapus",
	})
}
