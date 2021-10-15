package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch1 := createAndFillChannel("First")
	ch2 := createAndFillChannel("Second")
	finalChannel := converge(ch1, ch2)

	for i := 0; i < 16; i++ {

		fmt.Println(<-finalChannel)
	}

	close(ch1)
	close(ch2)
	close(finalChannel)
}

func createAndFillChannel(value string) chan string {

	ch := make(chan string)

	go func(ch chan string, valueA string) {

		// Another infinity loop
		for i := 0; ; i++ {

			ch <- fmt.Sprintf("Func %v, says: %v", value, i)

			// the code bellow results into 0 and 1 milisecond
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}

	}(ch, value)

	return ch
}

func converge(x, y chan string) chan string {

	newChan := make(chan string)

	go func() {

		for {

			// newChan will receive everything from channel x
			newChan <- <-x
		}

	}()

	go func() {

		for {

			// newChan will receive everything from channel y
			newChan <- <-y
		}
	}()

	return newChan
}
