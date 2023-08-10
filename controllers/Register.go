package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"yoga/task-a/entities"
)

func Register(db *sql.DB) {

	Register := entities.Users{}

	fmt.Print("Masukkan nama anda: ")
	fmt.Scanln(&Register.Name)

	fmt.Print("Masukkan email anda: ")
	fmt.Scanln(&Register.Email)

	fmt.Print("Masukkan password anda: ")
	fmt.Scanln(&Register.Password)

	fmt.Print("Masukkan alamat anda: ")
	fmt.Scanln(&Register.Address)

	fmt.Print("Masukkan nomor telfon anda: ")
	fmt.Scanln(&Register.Phone_number)

	result, errInsert := db.Exec("INSERT INTO users (name, email, password, address, phone_number) values (?, ?, ?, ?,?)", Register.Name, Register.Email, Register.Password, Register.Address, Register.Phone_number)
	if errInsert != nil {
		log.Fatal("error insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Berhasil mendaftar")
			MenuPilihan(db)
		} else {
			fmt.Println("failed to insert data")
		}
	}
}
