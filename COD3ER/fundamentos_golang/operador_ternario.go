package main

import "fmt"

//Não existe operador ternário em Go
func main() {
	resultado := obterResultado(10)
	fmt.Println(resultado)
}

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}
