package main

//Pacote é um conjunto de arquivos que está na mesma pasta.
//Se desejar criar mais de um pacote, é preciso criar mais de uma pasta.

//O professor criou o módulo com "go mod init <NOME_DO_MODULO>"
//Preciso entender exatamente para que serve a criação de um módulo.
//Será que não funcionaria tudo sem dar o "go build" e sem criar esse init de modulo?

//Ao usar "go install", o arquivo buildado é instalado dentro da pasta principal do Go

//Para o arquivo ser executado, é preciso ter um pacote chamado main

//Quando lidamos com mais de um pacote no mesmo projeto, é preciso criar um módulo.
//O módulo é um conjunto de pacotes criados no projeto (Um pacote de pacotes!).

import (
	"fmt"
	"modulo/auxiliar" //É preciso importar o outro pacote
)

func main() {
	fmt.Println("Testando...")
	auxiliar.Escrever()
}
