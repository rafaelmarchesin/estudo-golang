package main

import "fmt"

func main() {
	//Para array é passado o número de posições e o tipo dos valores do array
	//Todos os elementos devem ter o mesmo tipo
	var array1 [5]int
	array1[0] = 123
	fmt.Println(array1)

	array2 := [5]string{
		"Posição 1",
		"Posição 2",
		"Posição 3",
		"Posição 4",
		"Posição 5",
	}
	fmt.Println(array2)

	//No array, se tentarmos colocar mais valores do que o tamanho dele, vai dar erro

	//Esses ... fixam o tamanho do array com a quantidade de valores que eu inseri na declaração literal,
	//mas, depois que eu inseri os valores, ele vai limitar o array nesse tamanho que criado na declaração
	array3 := [...]int{
		1,
		2,
		3,
	}

	fmt.Println(array3)

	//O slice não tem limitação de tamanho
	//É mais usado do que array
	//O slice também é refem do tipo, ou seja todos os dados inseridos dentro do slice devem ser do mesmo tipo
	slice := []int{
		1,
		2,
		3,
		4,
		5,
		6,
	}

	fmt.Println(slice)

	//O slice, na verdade, não é um array, é um ponteiro para um array
	//Se o slice estiver referenciando um array específico, qualquer alteração no array aparecerá no slice

	//O append cria um novo slice acrescentando valores nas próximas posições livres do slice
	slice = append(slice, 18)

	//-----------------------------------
	//-------- ARRAYS INTERNOS ----------
	//-----------------------------------

	//Tem o tamanho 10 e a capacidade 15
	//O que é a capacidade?
	slice3 := make([]float32, 10, 15)
	//Essa função make criou um slice que referencia 10 posições de um array que possui 15 posições.

	fmt.Println(slice3)

	//Nesta caso, apenas referenciamos o tamanho do slice
	//Quando não passamos a capacidade, a capacidade do slice é o tamanho dele, neste caso, 4
	slice4 := make([]float32, 4)
	fmt.Println("Capacidade:", cap(slice4))
}
