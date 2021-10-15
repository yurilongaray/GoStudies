package main

import "fmt"

func f1(x, y int) (b, p int) {

	b = x
	p = y

	return //This will return b and p
}

func main() {

	fmt.Println(f1(3, 5))
}
