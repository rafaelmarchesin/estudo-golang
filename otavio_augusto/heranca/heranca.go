package main

import "fmt"

func main() {
	pessoa1 := pessoa{
		nome:  "João",
		idade: 55,
	}

	estudante1 := estudante{
		pessoa:    pessoa1, //Para receber aqui a pessoa, é precio ter criado uma pessoa antes
		curso:     "Engenharia",
		faculdade: "USP",
	}

	fmt.Println(pessoa1.idade, estudante1.idade)
}

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa    //Isso é a herança, nesse caso, não é preciso especificar o tipo
	curso     string
	faculdade string
}
