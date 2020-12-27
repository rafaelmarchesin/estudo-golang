package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	//O código abaixo não é possível, pois só é possível usar append com slice e não com o array diretamente
	//array1 = append(array1, 4, 5, 6)

	slice1 = append(slice1, 4, 5, 6)

	fmt.Println(array1, slice1)

	slice2 := make([]int, 2) //Criei um slice com tamanho 2
	copy(slice2, slice1)     //Copia os dois primeiros elementos do slice1 no slice2, já que o 2 é limitado a apenas 2 elementos. Ele não expande o tamanho do slice
	fmt.Println(slice2)      //O copy() também não pode passar um array nem mesmo na origem dos dados, os dois elementos devem ser slice
}
