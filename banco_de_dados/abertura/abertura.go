package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Ao colocar no início a "_", esse package é importado, mas ignorado, aí ele não é deletado quando salva
)

func main() {
	//Ao colocar essa string "mysql", chama o package mysql que importamos. Esse package funciona como um driver para realizar a conexão com o banco
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/") //O professor conecta sem esse "tcp", mas vi no stackoverflow que é preciso e funcionou
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	// exec(db, "drop table if exists usuarios")
	exec(db, `create table if not exists amigos(
		id integer auto_increment,
		nome varchar(90),
		PRIMARY KEY (id)
	)`)
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err) //O que será que faz essa função panic()?
	}
	return result
}
