package main

import (
	"fmt"      //pacote principal da aplicação
	"net/http" //pacote de requisições web
	"os"       //pacote do sistema operacional
)

func main() {

	introducao()
	for {
		escolha := menu()
		controllerComando(escolha)
	}

}

func introducao() {

	versao := 1.1
	var nome string
	fmt.Println("Qual o seu nome: ")
	fmt.Scan(&nome)

	fmt.Println("Olá senhor", nome+", o programa está na versão ", versao)
}

func menu() int {
	var escolha int

	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
	fmt.Scan(&escolha)

	return escolha
}

func controllerComando(escolha int) {

	switch escolha {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando inválido")
		os.Exit(-1)

	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	resposta, _ := http.Get("https://alura.com.br")

	if resposta.StatusCode == 200 {
		fmt.Println("Carregado com sucesso")
	} else {
		fmt.Println("Erro ", resposta.StatusCode)
	}

}
