package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func ReadAccount(db *sql.DB) {
	var Nama string
	fmt.Print("Masukan nama lengkap anda: ")
	fmt.Scanln(&Nama)

	rows, errSelect := db.Query("SELECT id, name, email, saldo, phone_number, address FROM users WHERE name = ?", Nama)

	if errSelect != nil {
		log.Fatal("error running select query: ", errSelect.Error())
	}

	defer rows.Close()

	found := false

	for rows.Next() {
		var id int
		var name string
		var email string
		var saldo string
		var phone_number string
		var alamat string

		err := rows.Scan(&id, &name, &email, &saldo, &phone_number, &alamat)
		if err != nil {
			log.Fatal("error scanning row: ", err.Error())
		}

		found = true

		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		fmt.Println("Saldo:", saldo)
		fmt.Println("Nomor telfon:", phone_number)
		fmt.Println("Alamat:", alamat)

		if !found {
			fmt.Println("Data tidak ditemukan.")
		}
	}
}
