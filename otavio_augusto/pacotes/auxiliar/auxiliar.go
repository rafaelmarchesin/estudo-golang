//O nome da pasta e do arquivo não importam. O que importa é o nome inserido após o "package"

package auxiliar

import "fmt"

//Escrever é uma função dentro do pacote "auxiliar"
func Escrever() { //Com letra maiúscula, a função pode ser exportada para fora do pacote
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
}
