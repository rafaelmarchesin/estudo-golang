package main

import "fmt"

//Este exercício tem como objetivo, provar que mais de um slice apontam para um mesmo array
func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[0] = 7 //Inseri o novo valor no primeiro slice

	fmt.Println(s1, s2)
}
