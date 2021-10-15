package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Convergence is to get things from many "ways" and put them all into a single "way"

func main() {

	pairChannel := make(chan int)
	oddChannel := make(chan int)
	converge := make(chan int)

	go sendVals(pairChannel, oddChannel)
	go receiveVals(pairChannel, oddChannel, converge)

	for v := range converge {

		fmt.Println(v)
	}

	fmt.Println("Converge is to pass all values from channels to a single channel and then make all logic in this channel (pairChannel + oddChannel = converge)")
}

func sendVals(pairChannel, oddChannel chan int) {

	const total = 10

	for i := 0; i < total; i++ {

		if i%2 == 0 {

			pairChannel <- i
		} else {

			oddChannel <- i
		}
	}

	close(pairChannel)
	close(oddChannel)
}

func receiveVals(pairChannel, oddChannel, converge chan int) {

	wg.Add(2)

	go func() {

		for v := range pairChannel {

			converge <- v
		}

		wg.Done()
	}()

	go func() {

		for v := range oddChannel {

			converge <- v
		}

		wg.Done()
	}()

	wg.Wait()

	close(converge)
}
