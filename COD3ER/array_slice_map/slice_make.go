package main

import "fmt"

func main() {
	s := make([]int, 10) //Cria um slice com 10 elementos
	s[9] = 12            //Atribui o valor 12 ao último espaço do slice
	fmt.Println(s)

	s = make([]int, 10, 20)        //Nesse caso, o slice tem 10 elementos, mas 20 espaços para o array interno
	fmt.Println(s, len(s), cap(s)) //cap() é a capacidade do array

	//Abaixo, usamos o append() para acrescentar mais 10 valores ao tamanho do slice e, assim, ocupar todo o espaço do array a que ele se refere
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	//Quando criamos o slice pelo "make" (não só com o "make"), ele consegue receber mais valores do que a capacidade máxima do array,
	//isso acontece porque o slice referencia outro array e dobra sua capacidade evitando, assim, erros
	s = append(s, 1, 2, 3, 4)
	fmt.Println(s)
	fmt.Println(s, len(s), cap(s))

	//Fiz esse teste para confirmar se o aumento do slice só acontece com o "make", mas acabei de ver que não
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:3]
	fmt.Println(array, slice)

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(slice)
}
