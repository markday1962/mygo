package main

import (
	"fmt"
)

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("about to exit")

}

func receive(e, o <-chan int, q <-chan bool) {
	//infinite loop, that pulls from the channels until a quit
	//is pulled from the channel
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel", v)
		case v := <-o:
			fmt.Println("from the odd channel", v)
		case i, ok := <-q:
			//comma ok idiom, to check a channel is closed
			if !ok {
				fmt.Println("from comma ok", i)
				return
			} else {
				fmt.Println("from comma ok", i)
			}
		}
	}
}

func send(e, o chan<- int, q chan<- bool) {
	//generate a bunch off values and put them on the channels
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
