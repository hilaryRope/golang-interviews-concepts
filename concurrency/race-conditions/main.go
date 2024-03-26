package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	counter := 0
	const gs = 100

	wg.Add(gs) // will add all go routines and wait for all

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()

	fmt.Println("Hello, playground")
	fmt.Println("the counter is", counter)

}
