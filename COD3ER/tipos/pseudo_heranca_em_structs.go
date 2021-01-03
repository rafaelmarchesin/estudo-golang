package main

import "fmt"

//O que é herança? O professor explicou um pouco, mas eu não entendi muito bem
//Pelo que eu entendi, herança seria quando uma struct recebe outra struct, mas o professor
//chamou de pseudo-herança, pois, na verdade, uma está dentro da outra e não se tornaram uma
func main() {
	f := ferrari{} //Foi instanciado apenas ferrari
	f.nome = "F40" //Mas pegou as informações de "carro", também. É como se expandisse ao criar um campo anônimo com o tipo da struct "carro"
	f.velocidadeAtual = 0
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
}

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       //campos anônimos
	turboLigado bool
}
