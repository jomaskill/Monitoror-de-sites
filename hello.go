package main

import (
	"bufio"    //pacote para abrir arquivos
	"fmt"      //pacote principal da aplicação
	"io"       //verificação de IOF
	"net/http" //pacote de requisições web
	"os"       //pacote do sistema operacional
	"strings"
)

func main() {
	introducao()

	//um for sem parametros, não possui critério de parada
	for {
		escolha := menu()
		controllerComando(escolha)
	}

}

func introducao() {
	//declaração de variavel normal 'var nomeDaVariavel Tipo'
	//declaração de variavel curta ' nomeDaVariavel := Valor'
	versao := 1.1
	var nome string

	fmt.Println("Qual o seu nome: ")
	fmt.Println("")
	fmt.Scan(&nome)
	fmt.Println("")
	fmt.Println("Olá senhor", nome+", o programa está na versão ", versao)
	fmt.Println("")
}

func menu() int {
	var escolha int

	fmt.Println("1 - Iniciar o monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
	fmt.Println("")
	fmt.Scan(&escolha)

	return escolha
}

//parametro da função 'Nome Tipo'
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

	//slices, são arrays que não precisam declaram o tamanho
	sites := sitesArquivo()

	//para cada 'sites' que existe, ele coloca na variavel 'item'
	//range retorna 'posicao, variavel'
	for _, item := range sites {
		fmt.Println("Monitorando...", item)

		//http.get retorna 'resposta, log'
		resposta, erro := http.Get(item)

		if erro != nil {
			fmt.Println("ocorreu um erro")
		}

		//sucesso na requisição
		if resposta.StatusCode == 200 {
			fmt.Println("Carregado com sucesso")
		} else {
			fmt.Println("Erro ", resposta.StatusCode)
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func sitesArquivo() []string {

	var sites []string

	//Abre arquivo
	arquivo, erro := os.Open("sites.txt")

	if erro != nil {
		fmt.Println("ocorreu um erro")
	}

	//salva no formato de possível leitura
	leitor := bufio.NewReader(arquivo)

	for {
		// lê até fim da linha
		linha, erro := leitor.ReadString('\n')
		//Tira o '/n' do arquivo
		linha = strings.TrimSpace(linha)

		//adiciona site ao array
		sites = append(sites, linha)

		//ao chegar no final do arquivo, sai
		if erro == io.EOF {
			break
		}

	}

	//Fecha arquivo
	arquivo.Close()
	return sites
}
