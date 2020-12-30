package main

import (
	"fmt"
	"time"
)

func main() {
	c := juntar(falar("joão"), falar("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	//fmt.Println(<-c, <-c) >> gera deadlock porque pedimos para retornar só 6 resultados
}

func falar(pessoa string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", pessoa, i) //toda vez que passar 1 sec, manda para o canal
		}
	}()
	return c
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-entrada1:
				c <- s
			case s := <-entrada2:
				c <- s
			}
		}
	}()
	return c
}
