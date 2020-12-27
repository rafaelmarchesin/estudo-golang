package main

import "fmt"

//Só pode ter um main() para cada package (pasta), mas pode ter um init() para cada arquivo
//Dentro de uma pasta, só um arquivo pode ter main()
func main() {
	fmt.Println("Main...")
}

//A função init é executada quando o arquido em Go é interpretado
//Essa função não precisa ser chamada no main()
func init() {
	fmt.Println("Iniciando...")
}
