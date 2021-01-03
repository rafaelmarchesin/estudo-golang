package main

import (
	"fmt"

	html "./pacote_reutilizavel"
)

func main() {
	c := juntar(
		html.Titulo("https://www.google.com", "https://www.amazon.com.br/"),
		html.Titulo("https://www.terra.com.br", "https://www.uol.com.br"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

func encaminhar(origem <-chan string, destino chan string) {
	//Quando tem valor em origem, insere em destino, quando não tem, não insere, mas fica o tempo todo funcionando esse "for"
	for {
		destino <- <-origem
	}
}

//Função responsável por multiplexar (é misturar as mensagens de fontes diferentes e juntar no mesmo canal)
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
