package controllers

import (
	"database/sql"
	"fmt"
)

func MenuPilihan(db *sql.DB) {

	fmt.Println("Silahkan input pilihan anda")
	fmt.Println("Pilih Menu:\n1. ReadAccount.\n2. UpdateAccount. \n3. Delete.\n4. TopUp. \n5. Transfer.\n6. HistoryTopUp. \n7. HistoryTransfer.\n8. CariUser. \n9. Logout")

	var pilihan int

	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		ReadAccount(db)

	case 2:
		UpdateAccount(db)

	case 3:
		DeleteAcc(db)

	case 4:
		TopUp(db)

	case 5:
		Transfer(db)

	case 6:
		HistoryTopUp(db)

	case 7:
		HistoryTransfer(db)

	case 8:
		CariUser(db)

	case 9:
		LogOut(db)

	}
}
