package main

import "fmt"

func main() {
	n := 1

	inc1(n)        //Neste caso, o valor de "n" não é alterado, pois é passado o valor e não o endereço de memória
	fmt.Println(n) //Aqui iríamos usar só o valor para fazer outros cálculos e guardar em outros locais de memória

	inc2(&n)       //Aqui o valor é alterado, pois passamos o endereço de memória e a alteração é feita diretamente no endereço
	fmt.Println(n) //Como passamos o endereço, agora, sim, a alteração consegue ser efetiva

}

func inc1(n int) {
	n++
}

//REVISÃO: Um ponteiro é representado por um "*"
func inc2(n *int) {
	//REVISÃO: "*" é usado para desreferenciar, ou seja,
	//ter acesso ao valor no qual o ponteiro aponta
	*n++
}

//Não é tão recomendável fazer isso, pois isso pode causar problemas para outras partes da programação
//O legal do Go é que ele prioriza passar o valor e isso evita efeitos colaterais
