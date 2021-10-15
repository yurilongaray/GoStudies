package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	// func (receiver) indentifier(parameter parameterType) (returnType) { code }

	fmt.Println(sum(10, 39))

	fmt.Println(showAll("Show this", 10, 39, 44))

	sliceTest := []int{1, 2, 3, 4}

	fmt.Println(showAll("LOL", sliceTest...))

	person1 := person{name: "Carlos", age: 99}

	person1.showPerson()

	// using defer says that the code in running at the function's end
	defer fmt.Println("Last")
	fmt.Println("First")

	// anonimous function = u declare and call it at the same time... and it is used a lot into go routines
	func(l int) {

		fmt.Println(l * 99)
	}(person1.age)

	// functions as expressions, function as a variable
	aFunction := func(bb int) int {

		return (bb * 10)
	}

	fmt.Println(aFunction(8))

	fmt.Println(daddy(288)(99))

	fmt.Println(callbackExample(sumAll, []int{1, 2, 3, 4, 5, 6}...))

	// calling closures
	closure1 := bb()

	fmt.Println("closure1", closure1())
	fmt.Println("closure1", closure1())
	fmt.Println("closure1", closure1())

	// calling the same closure but in another scope
	closure2 := bb()

	fmt.Println("closure2", closure2())
	fmt.Println("closure2", closure2())
	fmt.Println("closure2", closure2())

	func(b float64) {
		fmt.Println("This is an anonimous function", b)
	}(99.8)

	closure3 := i()

	fmt.Println("closure3", closure3())
	fmt.Println("closure3", closure3())
	fmt.Println("closure3", closure3())
}

func sum(x, y int) int {

	return x + y
}

// the parameter x (below) becomes a slice (parâmetro variádico)
func showAll(b string, x ...int) (string, int, int) {

	soma := 0

	for _, value := range x {

		soma += value
	}

	return b, soma, len(x)
}

// This is a method that is linked to person (person here is a receiver)
func (p person) showPerson() {

	fmt.Println(p.name)
}

// function returning function
func daddy(v int) func(int) string {

	return func(l int) string {

		return fmt.Sprintf("%v and %v", v, l)
	}
}

func sumAll(v ...int) int {

	var result int

	for _, value := range v {

		result += value
	}

	return result
}

// callback example
func callbackExample(f func(x ...int) int, y ...int) int {

	var oddValues []int

	for _, v := range y {

		if v%2 != 0 {

			oddValues = append(oddValues, v)
		}
	}

	total := f(oddValues...)

	return total
}

// Closure = function that returns another function with its scope
func bb() func() int {

	x := 0

	return func() int {

		x++
		return x
	}
}

func i() func() int {

	x := 0

	return func() int {

		x++

		return x
	}
}
