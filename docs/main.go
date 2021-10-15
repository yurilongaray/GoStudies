package main

import (
	"fmt"

	dog "exemple.com/hello/docs/pkg"
)

func main() {

	fmt.Println(dog.Age(22))
}

// This is a function from a pkg
func Other() {

	fmt.Println("This is from Other")
}
