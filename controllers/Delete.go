package controllers

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteAcc(db *sql.DB) {
	var deleteAkun string
	fmt.Print("Masukan nomor hp anda: ")
	fmt.Scanln(&deleteAkun)

	query := "DELETE FROM users WHERE phone_number = ?"
	result, errDelete := db.Exec(query, deleteAkun)
	if errDelete != nil {
		log.Fatal("error:", errDelete.Error())
	} else {
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal("error getting rows affected:", err.Error())
		}
		if rowsAffected > 0 {
			fmt.Println("Akun anda telah terhapus")
		} else {
			fmt.Println("No data deleted")
		}
	}
}
