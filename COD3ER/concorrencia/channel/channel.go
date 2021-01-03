package main

import "fmt"

//O channel é um tipo da linguagem como um float, um int
func main() {
	ch := make(chan int, 1) //foi criado um canal de inteiros. O segundo parâmetro é o buffer

	//Como fazemos para enviar uma informação para um canal?
	//Como fazemos para receber uma informação de um canal?

	ch <- 1 //Estamos enviando para o canal o valor 1 (ele recebe valores inteiros)
	<-ch    //Dessa forma, recebemos dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)

	//Neste exemplo, os processos são síncronos, mas, quando usarmos goroutines, teremos processos assíncronos.
}
