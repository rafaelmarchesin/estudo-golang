package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel() //Isso significa que esse teste pode rodar de forma paralela com outros testes
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64") //Não sei para que serve t.Skip
	}
	// ...
	t.Fail()
}
