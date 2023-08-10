package controllers

import (
	"database/sql"
	"fmt"
)

func Transfer(db *sql.DB) {
	var Pengirim string
	fmt.Println("Masukkan nomor telepon Anda:")
	fmt.Scanln(&Pengirim)

	var Penerima string
	fmt.Println("Masukkan nomor telepon tujuan:")
	fmt.Scanln(&Penerima)

	var Tf int
	fmt.Println("Masukkan nominal yang ingin Anda transfer:")
	fmt.Scanln(&Tf)

	var saldoAwalPengirim int
	err := db.QueryRow("SELECT saldo FROM users WHERE phone_number = ?", Pengirim).Scan(&saldoAwalPengirim)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	if saldoAwalPengirim < Tf {
		fmt.Println("Saldo Anda tidak mencukupi untuk transfer ini.")
		return
	}

	var saldoAwalPenerima int
	err = db.QueryRow("SELECT saldo FROM users WHERE phone_number = ?", Penerima).Scan(&saldoAwalPenerima)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	saldoAkhirPengirim := saldoAwalPengirim - Tf
	saldoAkhirPenerima := saldoAwalPenerima + Tf

	_, err = db.Exec("UPDATE users SET saldo = ? WHERE phone_number = ?", saldoAkhirPengirim, Pengirim)
	if err != nil {
		fmt.Println("Error updating sender's saldo:", err.Error())
		return
	}

	_, err = db.Exec("UPDATE users SET saldo = ? WHERE phone_number = ?", saldoAkhirPenerima, Penerima)
	if err != nil {
		fmt.Println("Error updating reciver's saldo:", err.Error())
		return
	}

	_, err = db.Exec("INSERT INTO transfers (user_sender, user_reciver, transaksi) VALUES ((SELECT id FROM users WHERE phone_number = ?), (SELECT id FROM users WHERE phone_number = ?), ?)", Pengirim, Penerima, Tf)
	if err != nil {
		fmt.Println("Error inserting transaction:", err.Error())
		return
	}

	fmt.Println("Transfer berhasil!\n", "Sisa saldo Anda:", saldoAkhirPengirim)
}
