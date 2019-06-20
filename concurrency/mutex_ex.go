package main

import (
	"fmt"
	"runtime"
	"sync"
)

//
// building on race_condition.go, the mutex functions allows for the
// blocking of a variable from other go routines until an operation has completed.
// and it used to prevent a race condition
// go run -race main.go
//
func main() {
	counter := 0
	const gs = 100
	fmt.Println("CPU's\t", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock() //the code is locked and another go routine cannot access the counter variable
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock() //the code is unlocked allowing access to the counter variable
			wg.Done()
		}()
		fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	fmt.Println("counter", counter) //the number is different each run

}
