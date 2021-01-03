package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogo")

	if err != nil {
		log.Fatal()
	}

	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Novo Nome", 1)
	stmt.Exec("Teste", 2)
}
