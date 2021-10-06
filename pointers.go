package main

import (
	"fmt"
)

type animal struct {
	name string
	age  int
}

// x is the var
// &x is the var's memory location
// *x is the value of the var's memory
// (*p).name changing an prop from a p pointer

func main() {

	x := 10
	y := &x

	fmt.Println(x)

	// changing x by y (that contains its memory path (pointer))
	*y++

	fmt.Println(x)
	fmt.Printf("%T, %T, %v, %v\n", &x, y, &x, *y)
	fmt.Println(x)

	increment(&x)

	fmt.Println(x)

	animal1 := animal{
		name: "Name1",
		age:  22,
	}

	fmt.Println(animal1)

	dontChangeMe(animal1)

	fmt.Println(animal1)

	changeMe(&animal1)

	fmt.Println(animal1)
}

func increment(x *int) {

	*x++

	fmt.Println(*x)
}

// You can only change a value by reference when using pointers
func dontChangeMe(p animal) {

	p.name = "New name"
	p.age = 99
}

func changeMe(p *animal) {

	// correct way
	(*p).name = "New name"
	p.age = 99
}
