package main

import (
	"fmt"
	"sync"
)

func main() {
	//bidirectional channels
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send2(even, odd)

	go receive2(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

//send to the directional in channel
func send2(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

//receive channel consisting of 2 in channels and 1 out channel
func receive2(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	// takes the values off two channels and put them on to one
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)

}
