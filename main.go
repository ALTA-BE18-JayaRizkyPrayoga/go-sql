package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tafa/task-b/controllers"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var connect_db = os.Getenv("MAIN_DB")
	fmt.Println("")
	db, err := sql.Open("mysql", connect_db)
	if err != nil {
		log.Fatal("error connettion to db", err.Error())
	}
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error", errPing.Error())
	} else {
		fmt.Println("")
	}
	defer db.Close()
	fmt.Println("Selamat Datang di Aplikasi Go.Trans")
	fmt.Println("Aplikasi yang digunakan untuk top up dan transfer uang ke sesama pengguna")
	fmt.Println("Silahkan input pilihan anda:")
	fmt.Println("Pilih Menu:\n1. Login.\n2. Register.")
	var pilihan int

	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		controllers.LoginUser(db)

	case 2:
		controllers.Register(db)

	}
}
