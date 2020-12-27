package main

import "fmt"

func main() {
	tipo := tipo(nil)

	fmt.Println(tipo)
}

//O switch realizado neste exercício avalia o tipo do parâmetro
//Ainda não sei dizer o que é o tipo "interface" (aparentemente é algo mais genérico)
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "número real"
	case func():
		return "função"
	case string:
		return "é uma string"
	default:
		return "Não sei qual tipo é"
	}
}
