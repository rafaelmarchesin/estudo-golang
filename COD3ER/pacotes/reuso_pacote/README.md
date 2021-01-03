Para criar um novo pacote, basta acessar home/<user>/go/src/gitjub.com/<nome_da_conta_github>/<nome_do_seu_pacote>.go

A programação do pacote segue o mesmo padrão de programação em Golang.

No exercício, criamos um pacote no GoPath com o nome area (arquivo area.go)

path: src/github.com/rafaelmarchesin/area/area.go

Código criado:

```
package area

import "math"

//Pi é uma proporção numérica definida pela relação entre
//o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

//Circ é uma função que calcula a circunferência a partir do raio
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a área do retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Esta não é uma função visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
```