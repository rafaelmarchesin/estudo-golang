//Muitos desenvolvedores escrevem primeiro o teste do que a função.
//Geralmente, quando criamos os testes primeiros, pensamos no máximo de problemas que podem quebrar a função
//O professor disse que é bastante válido começar com os testes

//Isso é uma abordagem orientada a testes

//Em Go, não se usa tanto o conceido de mock

//No teste unitário, a função é isolada ao máximo possível. Testes integrados chamam outras funções caso uma função dependa de outra

package matematica

import (
	"fmt"
	"strconv"
)

//Media é responsável por calcular a média de valores inseridos
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
