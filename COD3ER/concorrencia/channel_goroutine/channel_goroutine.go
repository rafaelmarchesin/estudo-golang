package main

import (
	"fmt"
	"time"
)

//Channel (canal): é a forma de comunicação entre goroutines, é o ponto de sincronismo para receber os dados assíncronos.
//Channel é um tipo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second) //Esse tempo de espera é só para percebermos a execussão do processo
	c <- 2 * base           //Envia dados para o canal

	//Neste caso, o processo só passa para esta parte quando a execussão do canal anterior terminar
	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c //Recebe os dados do canal
	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println(<-c) //Gera um erro, pois goroutine já terminou a execussão no canal
}
