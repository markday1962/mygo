package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//tcpConn()
	udpConn()
}

func udpConn() {
	//to simulate connecting to a logstash server run $nc -luk 1903
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1903", timeout)
	if err != nil {
		fmt.Println("Failed to connect to server")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular log message")
}

func tcpConn() {
	//to simulate connecting to a logstash server run $nc -lk 1902
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		fmt.Println("Failed to connect to server")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)
	logger.Println("This is a regular log message")
	logger.Panic("This is a panic")

	//Output from nc
	//example 2019/06/16 tonet.go:37: This is a regular log message
	//example 2019/06/16 tonet.go:38: This is a panic
}
