package main

import "fmt"

type car struct {
	name string
	year int
}

type vehicle struct {
	car      // anonimous property
	capacity int
}

func main() {

	vehicle1 := vehicle{}

	vehicle1.name = "test"
	vehicle1.capacity = 30

	fmt.Println("vehicle", vehicle1)
}
