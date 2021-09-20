// NEVER DO THIS
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

// execute this file with the command go run --race race-condition.go
func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	var x int64

	totalConcurrency := 1000

	wg.Add(totalConcurrency)

	for i := 0; i < totalConcurrency; i++ {

		go func() {

			atomic.AddInt64(&x, 1)

			runtime.Gosched()

			wg.Done()
		}()
	}

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Final Value", x)
}
