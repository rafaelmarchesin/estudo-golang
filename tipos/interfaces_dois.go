package main

import "fmt"

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1.turboLigado)

	var carro2 esportivo = &ferrari{"F40", false, 0} //Quando usamos a partir de interface, precisamos usar com ponteiros
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}
