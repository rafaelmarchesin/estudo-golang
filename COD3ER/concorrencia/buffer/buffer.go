package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go rotina(ch)

	fmt.Println(<-ch) //Segura até o ch <- 3, depois disso, libera a sequência de tarefas do programa principal
}

//Obrigatoriamente tem que esperar até o "ch <- 3", depois disso, está liberado para serem executados em concorrência
func rotina(ch chan int) {
	fmt.Println("Executou")
	ch <- 1
	ch <- 2
	ch <- 3 //Tudo o que estiver dentro dessa função até aqui será lido na espera >> "fmt.Println(<-ch)".
	ch <- 4 //A partir daqui, vai acompanhar o programa normal, se o programa terminar, essas linhas não serão lidas
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	ch <- 10
}
