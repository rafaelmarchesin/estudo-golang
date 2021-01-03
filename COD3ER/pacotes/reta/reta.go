package main

import "math"

//Inicializar um método com primeira letra maiúscula, temos um método público
//Um método público tem visibilidade dentro e fora do pacote

//Ao compilar o pacote, todos os arquivos serão compilados como se fossem um único só,
//por isso não podemos nos limitar ao arquivo e, sim, ao pacote
//No Go, não temos visibilidade limitada ao arquivo e, sim, ao pacote

//NomeMetodo -> é público
//nomeMetodo -> é privado

//--- começando o programa ---

//Quando criamos uma struct pública, somos obrigados a colocar um comentário antes dela

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distancia é responsável por calcular a distância linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
