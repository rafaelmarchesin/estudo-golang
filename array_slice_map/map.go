package main

import "fmt"

//O map é uma estrutura chave/valor
//A chave só aceita um único valor
func main() {
	//Importante: É preciso inicializar o "map"
	//var aprovados map[int]string --> esse modo não funciona com map

	aprovados := make(map[int]string) //O tipo dentro do colchetes é o tipo da chave e o fora é o do valor

	aprovados[12345678] = "Rafael Marchesin"
	aprovados[87654321] = "Ana Maria"
	aprovados[28374625] = "Pedro Augusto"

	fmt.Println(aprovados)

	fmt.Println("--------------")

	//Usando o range para imprimir os valores
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	//Imprimindo ou alterando um valor específico dentro do map
	aprovados[12345678] = "João Augusto"
	fmt.Println("--------------")
	fmt.Println(aprovados[12345678])

	//Deletando um registro específico
	delete(aprovados, 12345678)

	fmt.Println("--------------")
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
