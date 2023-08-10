package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"yoga/task-a/entities"
)

func HistoryTopUp(db *sql.DB) {
	var UserID int
	fmt.Print("Masukkan user ID Anda: ")
	fmt.Scanln(&UserID)

	rows, err := db.Query("SELECT id, user_id, transaksi FROM top_up where user_id = ?", UserID)
	if err != nil {
		log.Fatal("Query error: ", err)
	}

	var transactions []entities.Top_up
	for rows.Next() {
		var transaction entities.Top_up
		errScan := rows.Scan(&transaction.ID, &transaction.User_id, &transaction.Transaksi)
		if errScan != nil {
			log.Fatal("Scan error: ", errScan)
		}
		transactions = append(transactions, transaction)
	}

	for _, v := range transactions {
		fmt.Println(v.Transaksi)

	}
}
