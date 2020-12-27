package main

import "fmt"

//O que são Funções Variáticas?
//São funções que recebem quantidades de valores não definidos, por exemplo, podemos inserir dois números ou 10 números diferentes na função
func main() {
	//Podemos inserir a quantidade de valores que quisermos
	//Podemos deixar vazio, inclusive, nesse caso, não funciona, pois o length vai ser zero e dar erro na divisão
	fmt.Printf("A média é %.2f.\n", media(7.6, 10.0, 3.5, 6.2, 12.0))
	fmt.Printf("A média é %.2f.\n", media(1, 2, 3))
}

//Os três pontos indicam que podemos inserir quantos valores quisermos
//Os parâmetros são trabalhados como array
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}
