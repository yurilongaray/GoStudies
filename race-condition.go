// NEVER DO THIS
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

// execute this file with the command go run --race race-condition.go
func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	x := 0
	totalConcurrency := 1000

	wg.Add(totalConcurrency)

	for i := 0; i < totalConcurrency; i++ {

		go func() {

			y := x

			// same as time.sleep (yield)
			runtime.Gosched()

			y++
			x = y

			wg.Done()
		}()
	}

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Final Value", x)
}
