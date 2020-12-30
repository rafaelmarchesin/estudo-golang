package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) //Canal criado sem buffer

	go rotina(c)

	//Como a goroutine é executada em paralelo, não espera terminar de processar a função,
	//a impressão abaixo é executada imediatamente.
	//Mas, como chamamos o retorno do canal (<-c), nosso código irá esperar parar a goroutine antes de continuar

	//Dúvida: Posso criar várias goroutines e atribuir tudo a um mesmo canal? Assim o código só continuará depois que encerrar todas?

	fmt.Println(<-c)                   //Operação bloqueante
	fmt.Println("O dado já foi lido.") //Essa linha só é chamada depois que o dado foi lido
	fmt.Println(<-c)                   //Aqui gera um deadlock, pois todas as rotinas já foram executadas no processo anterior do canal
	fmt.Println("Fim")                 //Essa linha não será executada, pois ocorrerá um deadlock na linha anterior
}

//Quando um canal tem buffer, ele só bloqueia quando o buffer fica cheio.
//Como, neste exemplo, o canal não terá buffer, ele já bloqueará
func rotina(c chan int) {
	//Já entendi o conceido de channel e goroutines, só não entendi ainda como usar
	//O que está ocorrendo exatamente abaixo?
	time.Sleep(time.Second)
	c <- 1 //Operação bloqueada

	//Pelo que eu entendi, o "fmt.Println(<-c)" será executado quando chegar em "c <- 1".
	//Isso quer dizer que, a partir daqui, a goroutine continua sendo executada em conjunto com as linhas seguintes da aplicação.
	//Desse modo, nao temos como saber quem será executado primeiro, o print abaixo ou o 'fmt.Println("O dado já foi lido.")'

	fmt.Println("Isso foi executado só depois que foi lido no canal")
}

//O buffer de um canal vai recebendo as rotinas e elas são lidas de modo assíncrono até que se encha o buffer, aí, só entra
//novas tarefas quando o buffer for liberado.
//Um canal sem buffer, executa as tarefas de modo sequencial, ou seja, uma tarefa só é executada quando a outra tiver terminado de ser executada
