//Dentro da linguagem Go não é permitido realizar operações aritiméticas com ponteiros
//Um ponteiro nada mais é do que a referência ao espaço de memória
package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //o que o ponteiro armazena é o enderaço de memória.

	p = &i //&i faz referência apenas ao endereço da variável "i"

	//Apenas "p" chama o endereço
	//Colocar "*p", chamamos o valor armazenado no endereço

	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
	fmt.Println("i: ", i)
	fmt.Println("&i: ", &i)
}
