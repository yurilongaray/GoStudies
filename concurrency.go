package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	// shows how many cpus cores exists
	fmt.Println(runtime.NumCPU())

	wg.Add(2)

	fmt.Println(runtime.NumGoroutine())

	// this is called GoRoutines
	go iterate1()
	go iterate2()

	// this will be executer before the funcs above because they are going to be executed in the next thread
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func iterate1() {

	for x := 0; x < 100; x++ {

		fmt.Println("1°: ", x)

		time.Sleep(20)
	}

	wg.Done()
}

func iterate2() {

	for x := 0; x < 100; x++ {

		fmt.Println("2°: ", x)

		time.Sleep(20)
	}

	wg.Done()
}
