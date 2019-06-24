package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	fmt.Println("Starting")
	//go fanOutIn(c1, c2)
	go fanOutInThrottle(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Exiting")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v) //passes v into the go function
		wg.Wait()
		close(c2)
	}
}

func fanOutInThrottle(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
