package main

import "fmt"

func main() {
	imprimirResultado(7)
}

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}
