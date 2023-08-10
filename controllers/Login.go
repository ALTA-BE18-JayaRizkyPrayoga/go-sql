package controllers

import (
	"database/sql"
	"fmt"
	"yoga/task-a/entities"
)

func LoginUser(db *sql.DB) {
	var phone_number, password string
	fmt.Print("Masukan nomor telfon: ")
	fmt.Scanln(&phone_number)
	fmt.Print("Masukan password: ")
	fmt.Scanln(&password)

	var user entities.Users
	err := db.QueryRow("SELECT id, name, phone_number, password FROM users WHERE phone_number = ?", phone_number).Scan(&user.Id, &user.Name, &user.Phone_number, &user.Password)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	if user.Password == password {
		fmt.Println("Berhasil!")
		MenuPilihan(db)
	} else {
		fmt.Println("Log in gagal.")
	}
}
