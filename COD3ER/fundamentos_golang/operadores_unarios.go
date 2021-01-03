package main

import "fmt"

func main() {
	x := 1
	y := 2

	//Golang não tem incremento pré-fixado, possui apenas x++ (pós-fixado)

	x++

	fmt.Println(x)

	y--

	fmt.Println(y)
}
