package main

import (
	"errors"
	"fmt"
)

func main() {
	//Existem 4 tipos de int no Go
	//int8 (aguenta 8 bits), int16 (aguenta 16 bits), int32 (aguenta 32 bits) e int64 (aguenta 64 bits)

	var numero int8 = 10 //Se for muito grande, dá erro
	fmt.Println(numero)

	//Quando se coloca apenas int, ele usa a arquitetura do computador, se o PC for 32bits, setá int32, se for 64bits, será int64

	//O tipo "uint" é o int sem sinal
	var numero2 uint = 10 //Não aceita negativo
	fmt.Println(numero2)

	//Alias
	var numero3 rune = 1231 //"rune" é exatamente igual "int32"
	//Usamos "rune" quando estamos trabalhando com a tabela ASCII, isso é só um padrão de boas práticas

	var numero4 byte = 121 //"byte" é o mesmo que "uint8", portanto, não aceita negativo

	_, _ = numero3, numero4

	//Tipos float: podem ser "float32" ou "float64"
	//Não pode declarar só "float"

	//var numero5 float = 12121.12 //ERRADO!!!!!
	numero6 := 12121.12 //Isso está certo! Recebe o float conforme a arquitetura do computador
	_ = numero6

	//String
	var str string = "Apenas um teste"
	_ = str

	char := 'B' //Quando colocamso aspas simples, retornará o número da tabela ASCII desse caractere
	//No Go, não existe "char", essa variável acabou sendo do tipo "int"
	//Com aspas simples, só é permitido colocar um caractere, se colocar mais de um, dá erro
	fmt.Println(char)

	//Valor Zero
	//O valor zero é quando a variável é inicializada, mas não recebe nenhum valor

	var texto string
	var num int32
	fmt.Println(texto, "\n", num) //Retornará os valores Zero para cada tipo

	//Booleano
	//True e false
	var booleano bool = true
	_ = booleano

	//Tipo Erro
	//No Go, tratamento de erro é diferente de outras linguagens
	var erro error
	fmt.Println("O valor zero do tipo error é", erro)

	//Como atribuir valor de erro para o erro?
	//Para criar erro, é preciso usar um pacote "errors".
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)

	//O tipo erro é muito importante no Go
}
