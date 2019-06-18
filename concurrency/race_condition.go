package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("GOroutine", runtime.NumGoroutine())

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GOroutine", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GOroutine", runtime.NumGoroutine())
	fmt.Println("counter", counter)

}
