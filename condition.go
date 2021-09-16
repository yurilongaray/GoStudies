package main

import "fmt"

func main() {

	x := 60

	if x > 100 || x == 20 {

		fmt.Println("IF")
	} else if x < 100 && x > 50 && x != 99 {

		fmt.Println("ELSE IF")
	} else {

		fmt.Println("ELSE")
	}

	switch { // same using switch true {

	case x > 100:
		fmt.Println("CASE 1")

	case x < 100 && x > 50:
		fmt.Println("CASE 2")
		fallthrough // Says that can enter in the next conditions

	case x > 1:
		fmt.Println("CASE 3")

	default:
		fmt.Println("DEFAULT")

	}

	switch x {
	case 1, 2, 60:
		fmt.Println("X is in group 1")
	case 4, 5:
		fmt.Println("X is in group 2")

	}

	var unknwon interface{}

	unknwon = "22" // change to another value and see the type in the switch below

	switch unknwon.(type) {

	case bool:
		fmt.Println(("bool"))
	case float64:
		fmt.Println(("float64"))
	case int:
		fmt.Println(("int"))
	case string:
		fmt.Println(("string"))
	default:
		fmt.Println("Unknown")
	}
}
