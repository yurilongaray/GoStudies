package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	runManyGoRoutines(100)

	wg.Wait()
}

func runManyGoRoutines(quant int) {

	wg.Add(quant)

	for i := 0; i < quant; i++ {

		go func(x int) {

			fmt.Println("Current: ", x)

			wg.Done()
		}(i)
	}
}
