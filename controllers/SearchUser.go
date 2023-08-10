package controllers

import (
	"database/sql"
	"fmt"
	"tafa/task-b/entities"
)

func CariUser(db *sql.DB) {

	var Phone_number string
	fmt.Println("Masukan nomor telfon tujuan:")
	fmt.Scanln(&Phone_number)

	rows, err := db.Query("select id, name, email, address from users where  phone_number= ?", Phone_number)
	if err != nil {
		fmt.Println("query error ", err.Error())
	}

	var transaction []entities.Users
	for rows.Next() {
		var row entities.Users
		errScan := rows.Scan(&row.Id, &row.Name, &row.Email, &row.Address)
		if errScan != nil {
			fmt.Println("Scan error ", errScan.Error())
		} else {
			transaction = append(transaction, row)
		}
	}
	for _, v := range transaction {

		fmt.Println("ID:", v.Id)
		fmt.Println("Name:", v.Name)
		fmt.Println("Email:", v.Email)
		fmt.Println("Alamat:", v.Address)

	}

}
