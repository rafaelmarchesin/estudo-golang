package main

import "fmt"

func main() {
	retorno := selectNumber(10)
	fmt.Println(retorno)
}

func selectNumber(num int) string {
	switch num {
	case 1:
		return "1"
	case 10:
		return "10"
	default:
		return "Nada legal isso!"
	}
}
