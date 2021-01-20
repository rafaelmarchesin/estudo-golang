package main

import "fmt"

func main() {
	if num := 15; num > 15 {
		fmt.Println("Maio que 15")
	} else if num == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//Quando se cria a variável dentro do escopo, ela fica restrita a esse escopo
	if num2 := 16; num2 > 0 {
		fmt.Println("Maior que zero!")
	}
}
