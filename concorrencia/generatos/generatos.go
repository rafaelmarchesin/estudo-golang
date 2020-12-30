package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Estudo sobre padrões de concorrência em Golang

//<-chan - Esse é um canal só de leitura, não se passa dados para ele
//Na função main, não foi preciso criar um channel, nem chamar goroutines, já recebemos o canal pronto
func main() {
	t1 := titulo("https://www.rafaelmarchesin.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.terra.com.br")

	//Imprimir o resultado que será lido primeiro em cada inserção
	fmt.Println("Chegou primeiro - t1:", <-t1, "| t2:", <-t2)
	fmt.Println("Chegou em segundo - t1:", <-t1, "| t2:", <-t2)
	//fmt.Println("Chegou em terceiro - t2:", <-t2)
	//fmt.Println("Chegou em quarto - t2:", <-t2)
}

//Esse código um channel já preparado para executar as goroutines, não é preciso fazer nenhuma inserção a mais
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}
