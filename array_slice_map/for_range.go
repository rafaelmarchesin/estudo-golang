package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //Quem contará o tamanho do array, nesse caso, é o compilador. Podemos ver que inserimos 5 valores

	//Retornará o índice e os elementos de dentro do array
	for i, num := range numeros {
		fmt.Printf("%d) %d\n", i, num)
	}

	//Podemos ignorar o índice e pegar só os valores do array
	for _, num := range numeros {
		fmt.Println(num)
	}
}
