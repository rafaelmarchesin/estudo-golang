//Definitivamente eu não entendi essa aula!

package main

import (
	"fmt"
	"time"
)

//Dá para usar laço for com os channels, desse modo, podemos receber os dados de forma um pouco mais pausada

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c) //cap() é a capacidade do canal, ou seja, 30
	for primo := range c {
		fmt.Printf("%d", primo)
	}
	fmt.Println("Fim")
}

//Esta função define se o número é primo
func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		//Se entra no if, o número não é primo
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(qtd int, c chan int) {
	inicio := 2
	for i := 0; i < qtd; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}
