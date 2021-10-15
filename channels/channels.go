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

	channel := make(chan int)    //bidirecional
	channel1 := make(chan<- int) //sender
	channel2 := make(<-chan int) //receiver

	// We can convert bidirecional channels to senders or receivers
	// but we cant convert senders or receivers to bidirecional

	fmt.Printf("%v, %T\n", channel, channel)
	fmt.Printf("%v, %T\n", channel1, channel1)
	fmt.Printf("%v, %T\n", channel2, channel2)

	// converting bidirecional channel to senders or receivers
	fmt.Printf("channel is %T now\n", (chan<- int)(channel))
	fmt.Printf("channel is %T now\n", (<-chan int)(channel))

	go func() {

		channel <- 42
	}()

	fmt.Println(<-channel)

	orThis()

	// In GO Channels is not possible to have a funcion that send and receives at the same time so we put go in the first one to make it run after the other
	go send(channel)
	receive(channel)
}

// channels can be bidirectional or can only receive or only send
// sendo to channel: chan<-
// receive from channel: <-chan
// to message between channels, funcions need to be go routines (to generate concurrency)

func orThis() {

	// This is a Buffer (avoid this)
	channel := make(chan int, 1)

	channel <- 42

	fmt.Println(<-channel)
}

func send(s chan<- int) {

	s <- 42
}

func receive(r <-chan int) {

	fmt.Println("Value receive from channel is: ", <-r)
}
