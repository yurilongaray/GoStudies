package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	ca1 := make(chan int)
	ca2 := make(chan int)
	const funcs = 5

	go sendSomeNumbers(100, ca1)
	go loopChannel(funcs, ca1, ca2)

	for v := range ca2 {

		fmt.Println("CA2: ", v)
	}
}

func sendSomeNumbers(n int, ch chan int) {

	for i := 0; i < n; i++ {

		ch <- i
	}

	close(ch)
}

func loopChannel(funcs int, ch1, ch2 chan int) {

	fmt.Println("len(ch1): ", len(ch1)) // Channel doesnt have a length

	for i := 0; i < funcs; i++ {

		wg.Add(1)

		go func() {

			for v := range ch1 {

				ch2 <- work(v)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	close(ch2)
}

func work(x int) int {

	time.Sleep(time.Millisecond * 1000)

	return x
}
