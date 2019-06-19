package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//
// Package atomic provides low level atomic memory primitives to prevent race conditions
//
func main() {
	var counter int64 = 0 //int64 is required by package atomic
	const gs = 100
	fmt.Println("CPU's\t", runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1) //increment counter by 1
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter)) //read from counter
			wg.Done()
		}()
		fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Number of GOroutine running\t", runtime.NumGoroutine())
	fmt.Println("counter", counter) //the number is different each run

}
