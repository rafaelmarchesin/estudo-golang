package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//A variável "i" fica visível apenas dentro do bloco if..else
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!")
		fmt.Println(i)
	} else {
		fmt.Println("Perdeu!")
		fmt.Println(i)
	}

	fmt.Println(time.Now().UnixNano())
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	//fmt.Println("s: ", s)
	r := rand.New(s)
	//fmt.Println("r: ", r)
	return r.Intn(10)
}
