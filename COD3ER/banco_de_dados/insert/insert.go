package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogo")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)") //O valor será substituído no lugar da "?"
	stmt.Exec("Maria")                                            //Insere os dados na tabela
	stmt.Exec("João")
	res, _ := stmt.Exec("Rafael") //a função Exec() retorna alguns dados como ID, por exemplo, e um erro

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected() //Quantidade de linhas afetadas com a ação
	fmt.Println(linhas)
}
