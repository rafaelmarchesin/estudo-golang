package main

import "fmt"

func main() {
	f1()
	f2("valor 1", "valor 2")
	fmt.Println(f3())
	fmt.Println(f4("valor 1", "valor2"))
	teste1, teste2 := f5()
	fmt.Println(teste1, teste2)
	f6 := f6()
	fmt.Println(f6)
}

//É importante investir tempo para criar bons nomes de funções e variáveis
func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

//Se indicar que vai retornar uma string, é obrigatório retornar uma string
//"Se prometeu, tem que cumprir"
func f3() string {
	return "F3"
}

//Entra dois parâmetros e retorna uma única string
//Se os dois parâmetro de entrada são do mesmo tipo, podemos indicar uma única vez
func f4(p1, p2 string) string {
	f4 := "F4: " + p1 + " " + p2
	return f4
}

//Retornando múltiplos valores
func f5() (string, string) {
	return "F5: Retorno 1", "F5: Retorno 2"
}

//O professor faz um retorno semelhante a esse, mas comigo não funciona esse retorno, dá erro
//Entendi! Ele não usa Printf, ele usa Sprintf. Preciso descobrir qual a diferença entre as duas funções
func f6() string {
	return fmt.Sprintf("F6: Apenas um teste\n")
}
