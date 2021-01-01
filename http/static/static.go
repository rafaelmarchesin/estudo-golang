package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // Lê os arquivos da pasta ./public
	http.Handle("/", fs)                      //Quando tiver uma requisição na raiz da aplicação, passa esse handle

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil)) //Escuta o servidor nessa porta
}
