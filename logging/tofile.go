package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logfile, err := os.Create("/tmp/logfile.txt")
	if err != nil {
		fmt.Println("Error creating log file")
	}
	defer logfile.Close()

	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular message")
	logger.Fatalln("This is a fatal exit")
	logger.Println("This is a regular message but will not run because of Fatalln")
}
