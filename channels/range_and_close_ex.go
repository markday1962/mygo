package main

import (
	"fmt"
)

func main() {
	// make a channel of int
	c := make(chan int)

	//send to channel
	go func() {
		for i := 0; i < 100; i++ {
			c <- i //send to the channel
		}
		close(c)
	}()

	//range loop, pulls off a channel until it is closed, see line 16
	for v := range c {
		fmt.Println(v)
	}

}
