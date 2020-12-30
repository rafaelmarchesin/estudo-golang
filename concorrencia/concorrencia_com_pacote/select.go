package main

import (
	"fmt"
	"time"

	html "./pacote_reutilizavel"
)

//Select é uma estrutura para se trabalhar com concorrência dentro de Golang
//Select é muito parecido com switch
func main() {
	campeao := oSiteMaisRapido(
		"https://www.google.com",
		"https://www.amazon.com.br/",
		"https://www.terra.com.br",
	)
	fmt.Println(campeao)
}

func oSiteMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	//O site que der uma resposta mais rápida, será o primeiro a imprimir
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond): //A não ser que o timeoute seja mais rápido
		return "Todos perderam!"
		//O problema é que a execussão do default é a mais rápida de todas, por isso, precisamos comentar ele
		//default:
		//	return "Sem resposta ainda!"
	}
}
