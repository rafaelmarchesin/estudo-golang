package main

import "fmt"

type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	itens  []item //Significa que vai ter vários itens dentro desse pedido
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.12},
			item{3, 100, 3.12},
		},
	}
	fmt.Printf("O valor total do pedido é %.2f\n", pedido.valorTotal())
}
