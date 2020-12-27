package main

import "fmt"

//Forma de criar uma função anônima
var soma = func(a, b int) int {
	return a + b
}

//Existem casos em que uma função pode receber outra função como parâmetro
//E casos em que uma função retorna outra função
func main() {

	//Essas funções, apesar de estarem dentro de uma variável, são chamadas como funções normais
	fmt.Println(soma(2, 3))

	//Podemos ter uma função criada dentro de uma variável dentro da função main()
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2, 3))
}
