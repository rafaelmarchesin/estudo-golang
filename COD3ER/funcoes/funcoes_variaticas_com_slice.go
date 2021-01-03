package main

import "fmt"

func main() {
	//Definimos um slice
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}

	//Como estamos passando um slice com vários valores, também precisamos colocar os três pontos
	//Não podemos inserir um array dentro de uma função variática, precisamos usar slices
	imprimirAprovados(aprovados...)
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
