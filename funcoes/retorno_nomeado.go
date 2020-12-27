package main

import "fmt"

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
}

//Como já criamos as novas variáveis no retorno, não precisamos colocar ":="
//A técnica abaixo é chamada de retorno limpo, pois não é passado nada no "return"
func troca(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2

	return
}
