package main

import (
	"fmt"
)

/* THIS IS GOINGO TO THROW DEADLOCK ERROR */
/* WHY? because one go routine must handle the channel of the other. The same go routine cant handle its own channel */
/*
func main() {

	channel := make(chan int)

	channel <- 42

	fmt.Println(<-channel)
}
*/
func main() {

	channel := make(chan int) //bidirecional

	go loopThis(30, channel)

	// channel received X values and can use a range on it
	for value := range channel {

		fmt.Println("Received from channel: ", value)
	}
}

func loopThis(t int, channel chan<- int) {

	for i := 0; i < t; i++ {

		channel <- i
	}

	close(channel)
}
