//O objetivo deste exercício é converter uma struc para json e vice-versa
package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`   //Foi feito um mapeamento para json e como esses campos aparecerão nele
	Nome  string   `json:"nome"` //Quando um identificador começa com letra maiúscula, ele é público
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//Struct para Json
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 1899.90,
		Tags: []string{
			"Promoção",
			"Eletrônico",
		},
	}

	//O json.Marshal converte em Json e retorna, também, um erro, por isso ignoramos o segundo retorno no exemplo abaixo
	p1Json, _ := json.Marshal(p1) //Esse método retorna um slice de byte, por isso temos que converter para string na hora de imprimir
	fmt.Println(string(p1Json))   //Para saber que é um slice de byte, basta colocar o mouse sobre o método e ver o que ele retorna

	//Json para Struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":"8.90","tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) //Primeiro temos que converter a string de json para slice de bytes
	fmt.Println(p2.Tags[1])
}
