//Apenas uma brincadeira

package main

import "fmt"

func main() {
	fmt.Println(apto())
}

type programador struct {
	Idade int
}

func apto() bool {
	dev := programador{29}
	if dev.Idade > 30 {
		return false
	}
	return true
}
