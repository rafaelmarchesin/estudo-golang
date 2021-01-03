package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//Interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Junior"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.90}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	//Podemos passar direto para o método imprimir. Nesse caso abaixo, não estamos definindo que esse conteúdo é do tipo "imprimivel"
	//mas ele segue a mesma estrutura do tipo "imprimivel", então é aceito, também
	imprimir(produto{"Calça Jeans", 79.90})

	//Abaixo estamos passando os dados do produto sem dizer que a variável é do tipo "imprimivel", mas vai aceitar mesmo assim no método "imprimir()"
	p1 := produto{"Calça Jeans", 79.90}
	imprimir(p1)

}
