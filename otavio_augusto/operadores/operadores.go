package main

import "fmt"

func main() {
	//Os principais operadores de Go são bem parecidos com os da maioria das linguagens de programação

	var num1 int16 = 10
	var num2 int32 = 25

	//Isso não pode ser feito, pois os tipos são diferentes
	//soma := num1 + num2

	_, _ = num1, num2

	//Operadores de atribuição
	//É o :=  e o "var nome tipo"

	//Operadores relacionais
	//Igual, Maior que, Menor que
	//Diferente, Maior ou igual, menor ou igual

	//Operadores lógicos
	//E (&&), ou(||), negação(!)

	//Operadores unários
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 12
	numero %= 3 //Módulo
	fmt.Println(numero)

	//No Go, não existe o préfixado (--numero)

	//Operador ternário
	//No Go não existe operador ternário
	//A premissa do Go é que exista apenas uma única maneira de se fazer as coisas...
	//Como podemos escrever o if, não existe, então o operador ternário

}
