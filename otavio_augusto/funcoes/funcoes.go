package main

import "fmt"

func main() {
	res := somar(10, 20)
	fmt.Println(res)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	teste := f("Legal")
	_ = teste

	v1, v2, _, v4, v5 := apenasUmTesteDeRetorno() //Quando coloca underline, o retorno é ignorado
	fmt.Println(v1, v2, v4, v5)
}

//As funções em Go podem ter parâmetros e retornos
func somar(n1, n2 int) int {
	return n1 + n2
}

//No Go, as funções aceitam mais de um retorno por função
//É muito usado e é ótimo para lidar com erros
func apenasUmTesteDeRetorno() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}
