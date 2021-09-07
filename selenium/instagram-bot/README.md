# Introdução

Este projeto foi desenvolvido para aprender e estudar as ferramentas de automação de processos na web.

Para que ele rode corratamente, é preciso instalar o driver do Chrome (chromedriver) e, sem seguida, instalar o Selenium Standalone.

## Passos
### Instalando o Chrome Driver:
1. Baixar o chromedriver a partir do seguinte link:
[https://chromedriver.storage.googleapis.com/index.html?path=92.0.4515.107/]

2. Descompactar e mover o arquivo "chromedriver" para a seguinte pasta:
```
/usr/local/bin
```

3. Dar permissão de execussão para o arquivo:
```
$ chmod +x chromedriver
```

### Instalando o servidor Selenium:
1. Baixar a versão 4 ou mais recente do Selenium:
[https://www.selenium.dev/downloads/]
* Obs.: O arquivo é ".jar" e precisa do Java instalado para poder rodar.

2. Executar o servidor do Selenium com o comando abaixo:
```
$ java -jar selenium-server-4.0.0-rc-1.jar --ext example.jar:dir standalone --port 4444
```
* Obs.: Alterar o nome do arquivo do servidor caso seja necessário.