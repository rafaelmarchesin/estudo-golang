package main

import "fmt"

//Neste exercício, vamos ver um map dentro de outro map
func main() {
	funcionariosPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":  15256.78,
			"Gustavo Pereira": 5562.45,
		},
		"J": {
			"jeferson Luis": 1212.12,
		},
		"P": {
			"Márcio Gomes": 55623.45,
		},
	}

	for letra, nome := range funcionariosPorLetra {
		fmt.Println(letra)

		for n, salario := range nome {
			fmt.Println(n, salario)
		}
	}
}
