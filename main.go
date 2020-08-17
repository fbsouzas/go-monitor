package main

import (
	"fmt"
	"os"
)

func main() {
	showWelcomeMessage()
	showMenu()

	command := scanCommand()

	switch command {
	case 1:
		fmt.Println("Iniciando monitoramente")
	case 2:
		fmt.Println("Exigibindo logs")
	case 0:
		exitApp()
	default:
		fmt.Println("Comando não encontrado")
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
