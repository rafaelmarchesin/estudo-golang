package main

import "fmt"

type usuario struct {
	Nome     string
	Idade    uint8
	Endereco endereco
}

type endereco struct {
	Rua    string
	Numero uint8
}

func main() {
	fmt.Println("Estudo Structs")

	var u usuario
	fmt.Println(u) //Uma struct também tem um valor zero -> é o valor zero de cada tipo dentro da struct

	//Maneira 1 de popular
	u.Nome = "Marco"
	u.Idade = 22
	fmt.Println(u)

	//Maneira 2 de popular
	u2 := usuario{
		Nome:  "Rafael",
		Idade: 32,
	}
	fmt.Println(u2)

	//Maneira 3 de popular
	u3 := usuario{
		"Jonas",
		44,
		endereco{
			"Rua Nelson Solano",
			33,
		},
	}
	fmt.Println(u3)

	//Podemos ter uma struct dentro de uma struct
	u4 := usuario{
		Nome:  "Niwton",
		Idade: 69,
		Endereco: endereco{
			Rua:    "Lazuli",
			Numero: 12,
		},
	}
	fmt.Println(u4)
}
