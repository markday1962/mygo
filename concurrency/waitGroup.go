package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("Arch\t", runtime.GOARCH)
	fmt.Println("CPU's\t", runtime.NumCPU())
	fmt.Println("Goroutines running\t", runtime.NumGoroutine())
	fmt.Println("\n\n")

	//run a go routine
	wg.Add(1) //wait for one thing to complete in  our case foo()
	go foo()
	fmt.Println("Goroutines running\t", runtime.NumGoroutine())

	bar()
	fmt.Println("Goroutines running\t", runtime.NumGoroutine(), "\n")
	wg.Wait() //wait until wg.Done() in foo() is executed
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
