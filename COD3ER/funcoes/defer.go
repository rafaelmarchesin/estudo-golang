package main

import "fmt"

//Defer é uma espera até terminar algo antes de retornar a função
//É muito usado quando queremos fechar um recurso que abrimos antes de sair da função. Um exemplo pode ser uma conexão com o banco de dados.
func main() {
	fmt.Println(obterValorAprovado(6000))
}

//A sentença de código é deixada para o final com o "defer" mesmo que ele seja escrito no começo
func obterValorAprovado(num int) int {
	defer fmt.Println("Fim!") //Deixa para imprimir no final, só

	if num > 5000 {
		fmt.Println("Valor alto.")
		return 5000
	}

	fmt.Println("Valor baixo.")
	return num
}
