package main

import "runtime/debug"

//Uma pilha de funções (ou Stack de funções) é a ordem de execussão de processos durante o processamento
func main() {
	f1()
}

//A função executa primeiro, a função mais interna. Ela entra primeiro em "main()", passa por todas as funções e depois executa do final para o início
func f3() {
	debug.PrintStack() //Mostra a sequência de processos
}

func f2() {
	f3()
}

func f1() {
	f2()
}
