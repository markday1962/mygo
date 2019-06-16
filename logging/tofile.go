package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./logfile.txt")
	defer logfile.Close()

	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)
	logger.Println("This is a regular message")
	logger.Fatalln("This is a fatal exit")
	logger.Println("This is a regular message but will not run because of Fatalln")
}
