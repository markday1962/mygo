package main

import "fmt"

//channels block
func main() {
	//buffer channels allow values to sit in the channel until they are ready to be pulled off
	sbuffer()
	ubuffer()
}

func sbuffer() {
	//successful buffer, where a buffer that can take on value is created
	//buffer channels allow values to sit in the channel until they are ready to be pulled off
	c := make(chan int, 1) //make a channel which you can put on values of type int

	c <- 42                            //add 42 to the channel
	fmt.Println("Inside sbuffer", <-c) //take values of channel c
}

func ubuffer() {
	//unsuccessful buffer, where a buffer of 1 has been created but 2 values are trying
	//to be added to the buffer.
	c := make(chan int, 1) //make a channel which you can put on values of type int

	c <- 42                            //add 42 to the channel buffer
	c <- 43                            //add 43 to the channel buffer that nothing has been pulled off
	fmt.Println("Inside ubuffer", <-c) //take values of channel c

	// fatal error: all goroutines are asleep - deadlock!
	// goroutine 1 [chan send]:
	// main.ubuffer()
	// Users/markday/go/src/mygo/channels/buffer_channel_ex.go:27 +0x99
	// main.main()
	// Users/markday/go/src/mygo/channels/buffer_channel_ex.go:9 +0x25
}
