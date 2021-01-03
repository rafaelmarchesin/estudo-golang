package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

//Se não colocar o ponteiro, ele vai retornar para p1 sempre o primeiro nome inserido,
//ele não vai fazer a alteração efetivamente
func (p *pessoa) setNomeCompleto(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{nome: "Pedro", sobrenome: "Silva"}
	fmt.Println(p1.getNomeCompleto())

	p1.setNomeCompleto("Maria Aparecida")
	fmt.Println(p1.getNomeCompleto())
}
