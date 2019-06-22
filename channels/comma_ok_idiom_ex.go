package main

import (
	"fmt"
)

func main() {
	//make a channel to hold 3 values of type int
	c := make(chan int, 3)
	//send to the channel
	go func() {
		c <- 42
		c <- 43
		close(c)
	}()

	//receive from the channel
	v, ok := <-c
	fmt.Println(v, ok)
	//42 true, as a value has been pulled and the channel is open
	v, ok = <-c
	fmt.Println(v, ok)
	//43 true, as the next value has been pulled and the channel is open
	v, ok = <-c
	fmt.Println(v, ok)
	//0 false,  as the value has been pulled from the channel
	//and it has been closed

}
