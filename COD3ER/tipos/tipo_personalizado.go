package main

import "fmt"

type nota float64

//O professor disse que isso é criar um apelido (alias) para o reciver
func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

//Aqui podemos ver que chamamos o método usando apenas "n"
//Poderíamos, assim, associar vários métodos a esse tipo criado pela gente como se fosse um grupo de métodos
func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(3.4))
}
