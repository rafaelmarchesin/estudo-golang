package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1 //A variavel2 está em outro endereço de memória, o que fizemos foi apenas uma cópia do valor.

	//Como apenas copiamos o valor da variavel1 para a variavel2, o valor vopiado na variavel2 estará em um novo endereço de memória
	//Para que a variavel2 receba os valores alterados da variavel1, é preciso passar o ponteiro, ou seja, toda vez que tiver uma
	//alteração na variavel1, ele vai escutar e receber.

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//Ponteiro é uma referência de memória
	//O ponteiro indica o endereço da memória

	var variavel3 int
	var ponteiro *int
	var ponteiro2 *int

	variavel3 = 100
	ponteiro = &variavel3
	ponteiro2 = &variavel1

	fmt.Println(variavel3, ponteiro, *ponteiro, *ponteiro2) //Quando não colocamos o asterisco, recebemos o endereço de memória e não o valor armazenado nele.
}
