package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const numberTimesToTest = 3
const delay = 5
const applicationsFile = "applications.txt"
const logFile = "log.txt"

func main() {
	showWelcomeMessage()

	for {
		showMenu()

		command := scanCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			printLog()
		case 0:
			exitApp()
		default:
			fmt.Println("Command not found.")
		}
	}
}

func showWelcomeMessage() {
	fmt.Println("Welcome to Website Monitor")
}

func showMenu() {
	fmt.Println("")
	fmt.Println("Choose a command:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - View logs")
	fmt.Println("0 - Exit")
}

func scanCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("The chosen command was:", command)

	return command
}

func exitApp() {
	fmt.Println("Exiting of the Website Monitor")

	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Starting monitoring...")

	websites := readFile()

	for i := 0; i < numberTimesToTest; i++ {
		for _, website := range websites {
			websiteMonitors(website)
		}

		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
}

func websiteMonitors(website string) {
	resp, err := http.Get(website)

	if err != nil {
		fmt.Println(err)
	}

	message := formatMessage(website, resp.Status, resp.StatusCode == 200)

	writeLog(message)

	fmt.Print(message)
}

func readFile() []string {
	var websites []string

	file, err := os.Open(applicationsFile)

	if err != nil {
		fmt.Println(err)
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

func formatMessage(website string, httpStatus string, status bool) string {
	statusMessage := "RUNNING"

	if !status {
		statusMessage = "FAILURE"
	}

	return time.Now().Format("2006-01-02 15:04:05") +
		" - " +
		"[" + statusMessage + "]" +
		"[" + website + "]: " +
		httpStatus +
		"\n"
}

func writeLog(message string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(message)

	file.Close()
}

func printLog() {
	fmt.Println("Viewing logs...")

	file, err := ioutil.ReadFile(logFile)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))
}
