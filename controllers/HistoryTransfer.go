package controllers

import (
	"database/sql"
	"fmt"
	"tafa/task-b/entities"
)

func HistoryTransfer(db *sql.DB) {
	var UserID int
	fmt.Print("Masukkan user ID Anda: ")
	fmt.Scanln(&UserID)

	rows, err := db.Query("select id, user_sender, transaksi from transfers where  user_sender= ?", UserID)
	if err != nil {
		fmt.Println("query error ", err.Error())
	}

	var transaction []entities.Transfer
	for rows.Next() {
		var row entities.Transfer
		errScan := rows.Scan(&row.ID, &row.User_sender, &row.Transaksi)
		if errScan != nil {
			fmt.Println("Scan error ", errScan.Error())
		} else {
			transaction = append(transaction, row)
		}
	}
	for _, v := range transaction {
		fmt.Print(v.User_sender, "  ", v.Transaksi, "\n")
	}

}
