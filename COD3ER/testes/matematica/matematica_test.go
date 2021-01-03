package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado esperado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel() //Isso significa que esse teste pode rodar de forma paralela com outros testes
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
