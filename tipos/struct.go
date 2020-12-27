package main

import "fmt"

//Maneira de criar uma strict
type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//Receiver -> O que é?
//Pelo que entendi, é um método que tem uma conexão com a struct
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.99,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())
}
