package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/tfs")
	if err != nil {
		fmt.Println("C1")
		panic(err)
	}
	fmt.Println("Connect success")
	defer db.Close()
}