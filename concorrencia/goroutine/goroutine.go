package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//Processo de chamar as funções sem concorrência
	//fale("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "Porque só posso falar depois que você ficar quieta!", 1)

	//Conceito de thread é abrir uma linha de execussão independente
	//Uma única thread pode administrar inúmeras goroutines
	//Uma goroutine é chamada com o uso da expressão "go"
	//Um exemplo de thread é estar com alguém na fila do cinema, um compra o lanche e o outro compra os ingressos
	//É criar processos independentes (linhas de execussão independentes)
	//GoRoutines são funções que são executadas de forma independente
	go fale("Maria", "Por que você não fala comigo?", 500)
	go fale("João", "Porque só posso falar depois que você ficar quieta!", 500)

	//As vezes, a goroutine é mais devagar do que toda a sequência da nossa função main(), aí, o processamento do programa finaliza antes de terminar a goroutine
	//Nesse caso, não dá tempo de ter o retorno
	time.Sleep(time.Second * 5)
}
