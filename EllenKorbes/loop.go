package main

import "fmt"

// There is no WHILE in GO

func main() {

	for index := 0; index < 10; index++ {

		// fmt.Println("Index", index)
	}

	// Creating a Clock (Nested loops)
	for hour := 1; hour <= 12; hour++ {

		fmt.Println("\nHour:", hour)

		for minute := 1; minute <= 60; minute++ {

			// fmt.Printf("%v ", minute)
		}
	}

	for x := 33; x <= 122; x++ {

		fmt.Printf("%T, %v\n", x, string(x))
	}

	// https://web.fe.up.pt/~ee96100/projecto/Tabela%20ascii.htm
	for i := 65; i < 90; i++ {

		fmt.Println(string(i))

		for j := 0; j < 3; j++ {

			fmt.Printf("\t%#U\n", i)
		}
	}

	year := 1994
	current := 0

	for {

		if current == year {

			break
		}

		current++

		fmt.Print("Year: ", current, "; ")

		if current%2 == 0 {

			fmt.Print(current, " Is divisible by 2; ")
		}
	}
}
