package main

import "fmt"

//Uma função envolve outra função
//Nesse caso, ela só vai funcionar dentro da função em que ela está, não consegue funcionar do lado de fora
func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() //Vai executar o x de dentro da outra função, ou seja, vai retornar 10
}

func closure() func() {
	x := 10

	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}
