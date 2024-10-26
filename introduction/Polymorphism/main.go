package main

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
	fmt.Println("Console Logger: ", message)
}

type FileLogger struct {
	file *os.File
}

func (fl FileLogger) Log(message string) {
	if fl.file != nil {
		log.SetOutput(fl.file)
		log.Println("File Logger: ", message)
	} else {
		fmt.Println("File Logger: Error - File not opened")
	}
}

type RemoteLogger struct {
	Endpoint string
}

func (rl RemoteLogger) Log(message string) {
	fmt.Println("Remote Logger to %s: %s\n", rl.Endpoint, message)
}

func LogMessages(loggers []logger, message string) {
	for _, logger := range loggers {
		logger.Log(message)
	}
}

func main() {
	consolelogger := ConsoleLogger{}
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file: ", err)
		return
	}

	filelogger := FileLogger{file: file}

	remoteLogger := RemoteLogger{Endpoint: "http://logging.service"}

	loggers := []Logger{consolelogger, filelogger, remoteLogger}

	LogMessages(loggers, "This is a test log message.")

	defer file.Close()
}

/*
Explanation of the Example:
Interface Definition: The Logger interface defines a Log method that all logger types must implement.

Logger Types:

ConsoleLogger: Implements logging to the console.
FileLogger: Implements logging to a file. It uses os.File to handle file operations.
RemoteLogger: Implements logging to a remote service via an endpoint.
LogMessages Function: This function accepts a slice of Logger and iterates over each logger,
calling the Log method with the provided message.

Main Function:

Creates instances of ConsoleLogger, FileLogger, and RemoteLogger.
Stores them in a slice of Logger.
Calls LogMessages to log the same message using different logging strategies.
Output:
When you run this code, it will log to the console and append the message to the app.log file.
The remote logging will simulate sending the message to a remote service.

Benefits of Using Polymorphism in This Example:
Flexibility: New logger types can be added without modifying the existing logging logic.
For example, you could create an EmailLogger or a DatabaseLogger.
Interchangeability: The logging functionality is easily interchangeable.
You can use any logger type in the same context.
Single Responsibility: Each logger handles its specific responsibility,
adhering to the Single Responsibility Principle.

*/
