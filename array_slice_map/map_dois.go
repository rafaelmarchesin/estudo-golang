package main

import "fmt"

//Neste exercício, vamos inicializar o map de maneira literal
func main() {

	//Podemos declarar o mad direto na variável
	funccionariosESalarios := map[string]float64{
		"José João":        1332.23,
		"Marilia Gabriela": 12333.45,
		"Pedro Augusto":    325.12, //O último elemento também deve ter vírgula ou a chave encostada no elemento final "325.12}"
	}

	for nome, salario := range funccionariosESalarios {
		fmt.Println(nome, "=", salario)
	}
}
