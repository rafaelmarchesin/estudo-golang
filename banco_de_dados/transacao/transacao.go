//Uma transação significa que todos os comandos deverão ser executados ou nenhum comando será executado.package transacao
//Isso é importante para casos como "Apaga o dado de uma tabela", insere esse dado na outra tabela, mas obrigatoriamente, as duas tarefas devem ser executadas
//Se um ficar sem fazer está errado (exemplo, tirar o dinheiro de uma conta e aparecer na outra)

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(172.17.0.2)/cursogo")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	transaction, _ := db.Begin() //Inicia a transação
	stmt, _ := transaction.Prepare("insert into usuarios(id, nome) values(?,?)")
	stmt.Exec(2000, "Bianca")      //Como essa transação terá um erro, nenhum deles será inserido na tabela
	stmt.Exec(2001, "Jeferson")    //Nem Bianca e nem Jeferson
	_, err = stmt.Exec(1, "Tiago") //Chave duplicada para sumular um errado

	//Porém, para travar, obrigatoriamente temos que tratar o erro, senão é salvo os registro corretos e só o incorreto que não
	if err != nil {
		transaction.Rollback() //É a função que bloquia a execussão dos comandos inseridos
		log.Fatal(err)
	}

	transaction.Commit() //Executa o comando
}
