package strings

import (
	"strings"
	"testing"
)

//Essa é a mensagem padrão que será usada no teste
const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	t.Parallel() //Isso significa que esse teste pode rodar de forma paralela com outros testes
	//Esse parallel é executar usando concorrência para aumentar a perfórmance dos testes

	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Primeira String", "Primeira", 0}, //A segunda string tem que estar na posição 0 da primeira
		{"", "", 0},                        //A mesma coisa aqui
		{"Opa", "opa", -1},                 //Como não encontra, atual tem que receber "-1", se receber valor diferente, falha o teste
		{"Rafael", "f", 2},                 //"f" tem que estar na posição 2 de "Rafael"
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte) //Na verdade, estamos testando a função "strings.Index()"

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
