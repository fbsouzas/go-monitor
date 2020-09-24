# Go Monitor

This CLI project monitors the health check of your applications and generates logs with information about your applications is running or not.

## How it works

1 - Clone the project in your `src` dir.
```sh
git clone git@github.com:fbsouzas/go-monitor.git
```

2 - Go to the project dir.
```sh
cd go-monitor
```

3 - Create file with the applications for monitors.
```sh
cp applications.txt.sample applications.txt
```

4 - Puts your applications in the `applications.txt` file. Example:
```sh
https://www.my-aplication-1.com/
https://www.my-application2.com/
```

5 - Run the **Go monitor**
```sh
go run main.go
```

## Preview

## Goal

This project was created for studing and have the first contact with Golang and learn more about this programming language.

