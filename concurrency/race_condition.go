package main

import (
	"fmt"
	"runtime"
	"sync"
)

//
// Creating a go routine race condition, where each of the 100 go routine
// write and read to the variable counter
//
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
		fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	fmt.Println("counter", counter) //the number is different each run

}
