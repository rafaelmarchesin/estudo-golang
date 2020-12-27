package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //Array
	s1 := []int{1, 2, 3}  //Slice

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	//Slice não é um array, é um pedaço de um array
	//Não se cria um novo array, é apenas um ponteiro
	s2 := a2[1:3] //"s2" é uma fatia de "a2"

	s3 := a2[:4]

	s4 := s3[1:5] //Apesar de eu ter puchado um slice, ele pega ainda os valores do array, por isso consegui pegar 5 elementos

	fmt.Println(a2, s2, s3, s4)
}
