package main

import (
	"fmt"
)

func main() {

	pairChannel := make(chan int)
	oddChannel := make(chan int)
	quitChannel := make(chan bool)

	go sendNumbers(pairChannel, oddChannel, quitChannel)
	receiveValues(pairChannel, oddChannel, quitChannel)
}

func sendNumbers(pairChannel, oddChannel chan int, quitChannel chan bool) {

	const total = 100

	for i := 0; i < total; i++ {

		if i%2 == 0 {

			pairChannel <- i
		} else {

			oddChannel <- i
		}
	}

	close(pairChannel)
	close(oddChannel)

	quitChannel <- true
}

func receiveValues(pairChannel, oddChannel chan int, quitChannel chan bool) {

	for {

		select {

		// COMMA OK expression (same from struct but for channels)
		case v, ok := <-pairChannel:

			if !ok {

				fmt.Println("Not OK A!", v)
			} else {

				fmt.Println("A: ", v)
			}

		case v, ok := <-oddChannel:

			if !ok {

				fmt.Println("Not OK B!", v)
			} else {

				fmt.Println("B: ", v)
			}

		case v, ok := <-quitChannel:

			if !ok {

				fmt.Println("Not OK!", v)
			} else {

				fmt.Println("Finished!", v)
				return
			}
		}
	}
}
