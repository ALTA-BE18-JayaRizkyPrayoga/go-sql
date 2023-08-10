package controllers

import (
	"database/sql"
	"fmt"
)

func LogOut(db *sql.DB) {
	fmt.Println("Terimakasih sudah menggunakan layanan kami")
}
