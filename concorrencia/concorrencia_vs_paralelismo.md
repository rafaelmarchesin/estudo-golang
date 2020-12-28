# Concorrência vs Paralelismo

Concorrência: capacidade de administrar múltiplas tarefas. Isso ocorre num mesmo processador. O sistema operacional administra múltiplas tarefas em um único núcleo.

Paralelismo: fazer duas coisas exatamente ao mesmo tempo. Só é possível fazer isso no mundo dos computadores se tiver dois núcleos realizando processamento.

O paralelismo é mais custoso do que a concorrência. Ocorre o overhead adicional e uma tarefa feita em múltiplos processadores fica mais lenta (ocorre uma sobrecarga para que ocorra a divisão e a união dos dados entre os núcleos).

Go é uma linguagem concorrente. Mas lida muito bem com o paralelismo.

Paralelismo exige mais do hardware e do Sistema Operacional.

## Importante

Modelar o software imaginando a concorrência para que ele seja o máximo otimizado.

Ao planejar um software baseado em concorrência, precisamos pensar que ele será gerado a apartir da composição de processos (normalmente funções) que são executados de forma independente. Isso quer dizer que um processo, mesmo que seja interligado a outro, será executado separadamente de outro processo ligado a ele e, por isso, será necessário juntá-los no final. Em Go, fazer isso é bastante simples.