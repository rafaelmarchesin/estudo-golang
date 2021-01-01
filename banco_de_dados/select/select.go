package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogo")
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 1) //Retorna uma sequência de linhas
	defer rows.Close()                                                   //Acho que ele fecha o resultado do rows e aí ele vira só leitura, mas é só achismo mesmo

	for rows.Next() { //o Next() é usado para pegar todos os elementos até o último
		var u usuario
		rows.Scan(&u.id, &u.nome) //Já joga para dentro da struct
		fmt.Println(u)
	}
}
