package main

import (
	"fmt"
)

func main() {

	channelA := make(chan int)
	channelB := make(chan int)
	const b = 100

	go func(v int) {

		for i := 0; i < v; i++ {

			channelA <- i
			channelB <- i
		}
	}(b)

	for i := 0; i < b; i++ {

		select {
		case v := <-channelA:
			fmt.Println("A: ", v)
		case v := <-channelB:
			fmt.Println("B: ", v)
		}
	}

	fmt.Println("Finished!")
}
