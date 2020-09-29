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
	fmt.Println("Welcome to Go Monitor")
}

func showMenu() {
	fmt.Println("")
	fmt.Println("Choose a command:")
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - View logs")
	fmt.Println("0 - Exit")
}

func scanCommand() int {
	command := -1

	fmt.Scan(&command)
	fmt.Println("The chosen command was:", command)

	return command
}

func exitApp() {
	fmt.Println("Exiting of the Go Monitor")

	os.Exit(0)
}

func startMonitoring() {
	fmt.Println("Starting monitoring...")

	apps := readAppsFile()

	for i := 0; i < numberTimesToTest; i++ {
		for _, app := range apps {
			check(app)
		}

		fmt.Println("")

		time.Sleep(delay * time.Second)
	}
}

func check(app string) {
	resp, err := http.Get(app)

	if err != nil {
		fmt.Println(err)
	}

	message := formatLogMessage(app, resp.Status, resp.StatusCode == 200)

	writeLog(message)

	fmt.Print(message)
}

func readAppsFile() []string {
	var apps []string

	file, err := os.Open(applicationsFile)

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		apps = append(apps, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return apps
}

func formatLogMessage(app string, httpStatus string, status bool) string {
	statusMessage := "RUNNING"

	if !status {
		statusMessage = "FAILURE"
	}

	return time.Now().Format("2006-01-02 15:04:05") +
		" - " +
		"[" + statusMessage + "]" +
		"[" + app + "]: " +
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
