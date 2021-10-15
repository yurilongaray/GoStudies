// NEVER DO THIS
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

// execute this file with the command go run --race mutex.go
func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	x := 0
	totalConcurrency := 1000

	wg.Add(totalConcurrency)

	for i := 0; i < totalConcurrency; i++ {

		go func() {

			// locking with mutex
			mu.Lock()
			y := x

			runtime.Gosched()

			y++
			x = y
			// unlocking with mutex
			mu.Unlock()

			wg.Done()
		}()
	}

	fmt.Println("GoRoutines: ", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Final Value", x)
}
