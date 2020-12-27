package main

import (
	"fmt"
)

//OBS.: fatorial de negativo não existe
func main() {
	resultado := fatorial(10)
	fmt.Println(resultado)

	//Usar o "uint" faz já uma validação quando ao uso do sinal de negativo
	resultado2 := fatorial(-10)
	fmt.Println(resultado2)
}

//"uint" é um inteiro sem sinal
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}
