package main

import (
	"fmt"
	"net"
)

func main() {
	name := "www.aistemos.com"
	r, err := exists(name)
	if err != nil {
		fmt.Println("Error received:", err)
	}

	fmt.Printf("Connection Open for %v: %v",name, r)
}

func exists(name string) (bool, error){
	conn, err := net.Dial("tcp", name + ":80")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}
