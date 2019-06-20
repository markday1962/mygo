package main

import (
	"fmt"
)

func main() {
	c := make(chan<- int, 2) //send and receive channel
	cr := make(<-chan int)   //receive only channel
	cs := make(chan<- int)   //send only channel

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	stdChannel(42)
	sendChannel(42)

}

func stdChannel(i int) {
	c := make(chan int, 1)
	go func() {
		c <- i
	}()
	fmt.Println("Standard channel:", <-c) //pulling from the channel
}

func sendChannel(i int) {
	cs := make(chan<- int, 1) //sending an int to the channel
	go func() {
		cs <- i
	}()
	fmt.Println("Send Channel: {}", <-cs)
	//invalid operation: <-cs (receive from send-only type chan<- int)
}

func receiveChannel(i int) {
	cs := make(<-chan int, 1) //receiving from a chan of int
	go func() {
		cs <- i
		//invalid operation: cs <- i (send to receive-only type <-chan int)
	}()
	fmt.Println("Send Channel: {}", <-cs)
}
