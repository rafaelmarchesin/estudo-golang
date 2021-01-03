package main

import "fmt"

func main() {
	//Array possui o mesmo tipo de dados dentro dele
	//O array tem sempre o mesmo tamanho (não varia de tamanho)

	var notas [3]float64 //Foi definido o tamanho do array e o tipo dele.

	notas[0], notas[1], notas[2] = 4.5, 7.8, 10.0

	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))

	fmt.Printf("A média é %.2f\n", media)
}
