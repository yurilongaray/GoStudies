package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// https://www.youtube.com/watch?v=eoA4jVD7RQE&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=160&ab_channel=AprendaGo
func main() {

	achannel := make(chan int)
	const totalGoRoutines = 10

	go createGoRoutines(achannel, totalGoRoutines)
	getValuesFromChannel(achannel, totalGoRoutines)
}

func createGoRoutines(achannel chan int, totalGoRoutines int) {

	wg.Add(totalGoRoutines)

	for a := 0; a < totalGoRoutines; a++ {

		go func() {

			for i := 1; i <= 10; i++ {

				fmt.Printf("GoRouting (%v) \t Loop (%v)\n", a, i)
				achannel <- i * 10
			}

			wg.Done()
		}()
	}

	wg.Wait()

	close(achannel)
}

func getValuesFromChannel(ch chan int, totalGoRoutines int) {

	for v := range ch {

		fmt.Println("V: ", v)
	}
}
