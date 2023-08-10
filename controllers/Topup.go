package controllers

import (
	"database/sql"
	"fmt"
)

func TopUp(db *sql.DB) {
	fmt.Println("Top-up")

	var phoneToTopUp string
	fmt.Println("Masukan nomor telfon anda:")
	fmt.Scanln(&phoneToTopUp)

	var topUP int
	fmt.Println("Masukan nominal yang ingin anda top up:")
	fmt.Scanln(&topUP)

	var saldoAwal int
	err := db.QueryRow("SELECT saldo FROM users WHERE phone_number = ?", phoneToTopUp).Scan(&saldoAwal)
	if err != nil {
		fmt.Println("Error:", err.Error())

	}

	saldoAkhir := saldoAwal + topUP

	_, err = db.Exec("UPDATE users SET saldo = ? WHERE phone_number = ?", saldoAkhir, phoneToTopUp)
	if err != nil {
		fmt.Println("Error updating user's saldo:", err.Error())

	}

	_, err = db.Exec("INSERT INTO top_up (user_id, transaksi) VALUES ((SELECT id FROM users WHERE phone_number = ?), ?)", phoneToTopUp, topUP)
	if err != nil {
		fmt.Println("Error inserting top-up record:", err.Error())

	}

	fmt.Println("Top-up berhasil! Saldo anda:", saldoAkhir)

}
