package controllers

import (
	"database/sql"
	"fmt"
)

func UpdateAccount(db *sql.DB) {
	fmt.Println("Update data")

	var idToUpdate string
	fmt.Println("Input ID:")
	fmt.Scanln(&idToUpdate)

	fmt.Println("Input nama baru:")
	var newName string
	fmt.Scanln(&newName)
	fmt.Println("Input Email baru:")
	var newEmail string
	fmt.Scanln(&newEmail)
	fmt.Println("Input Password baru:")
	var newPassword string
	fmt.Scanln(&newPassword)
	fmt.Println("Input Address:")
	var newAddress string
	fmt.Scanln(&newAddress)
	fmt.Println("Input nomor telfon baru:")
	var newPhoneNumber string
	fmt.Scanln(&newPhoneNumber)

	_, err := db.Exec("UPDATE users SET name=?, email=?, password=?, address=?, phone_number=? WHERE id=?", newName, newEmail, newPassword, newAddress, newPhoneNumber, idToUpdate)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println("Data updated berhasil.")
	}
}
