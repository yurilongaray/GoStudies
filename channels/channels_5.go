package main

import (
	"fmt"
)

func main() {

	achannel := make(chan int)

	go sendSomeNumbers(achannel)
	retrieveValues(achannel)
}

func sendSomeNumbers(ch chan int) {

	for i := 1; i <= 100; i++ {

		ch <- i
	}

	close(ch)
}

func retrieveValues(ch chan int) {

	for v := range ch {

		fmt.Println("V: ", v)
	}
}
