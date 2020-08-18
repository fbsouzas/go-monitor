package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitories = 3
const delay = 5

func main() {
	showWelcomeMessage()

	for {
		showMenu()

		command := scanCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exigibindo logs")
		case 0:
			exitApp()
		default:
			fmt.Println("Comando não encontrado")
		}
	}
}

func showWelcomeMessage() {
	name := "Fábio"
	version := 0.1

	fmt.Println("Olá, Sr(a)", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exigibir os logs")
	fmt.Println("0 - Sair do programa")
}

func scanCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi:", command)

	return command
}

func exitApp() {
	fmt.Println("Saindo da aplicação")

	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Iniciando monitoramento...")

	websites := readerFile()

	for i := 0; i < monitories; i++ {
		for _, website := range websites {
			websiteTester(website)
		}

		time.Sleep(delay * time.Second)
	}
}

func websiteTester(website string) {
	resp, err := http.Get(website)
	message := "foi carregado com sucesso!"

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode != 200 {
		message = "está com problemas."
	}

	fmt.Println("O site", website, message, "Status code:", resp.StatusCode)
}

func readerFile() []string {
	var websites []string

	file, err := os.Open("websites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		websites = append(websites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return websites
}
