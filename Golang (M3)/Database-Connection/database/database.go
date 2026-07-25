package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error){
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/login_app")

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Database Berhasil Terkoneksi")
	return db, nil
}