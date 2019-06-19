package main

import "fmt"

//channels block
func main() {
	c := make(chan int) //make a channel which you can put on values of type int

	go func() {
		c <- 42 //add 42 to the channel
	}()
	fmt.Println(<-c) //take values of channel c
}
