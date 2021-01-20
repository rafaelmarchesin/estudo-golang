package main

import "fmt"

func main() {
	fmt.Println("--- MAPS ---")

	//O map é mais rígido, pois as chaves devem ser todas do mesmo tipo e o mesmo vale para os valores
	usuario := map[int]string{
		1: "Rafael",
		2: "Jorge",
	}

	fmt.Println(usuario[1])

	//Usando map como tipo de um map
	usuario2 := map[int]map[string]string{
		1: {
			"Nome":      "Rafael",
			"Sobrenome": "Marchesin",
		},
		2: {
			"Nome":  "Jorge",
			"Idade": "32",
		},
	}

	fmt.Println(usuario2[2])

	//Deletar uma chave
	delete(usuario2, 2)

	fmt.Println(usuario2)

	usuario2[3] = map[string]string{
		"Nome":   "John",
		"Altura": "1.13",
	}

	fmt.Println(usuario2)
}
