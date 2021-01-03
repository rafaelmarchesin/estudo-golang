package main

import "fmt"

func main() {
	retultado := exec(multiplicacao, 3, 6)
	fmt.Println(retultado)
}

func multiplicacao(a, b int) int {
	return a * b
}

//Neste caso, é passado 3 parâmetros: uma função que retorna um inteiro, p1 e p2
//A assinatura da função parâmetro tem que ser igual ao da função que irá entrar nessa nova função
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
