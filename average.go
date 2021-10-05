package main

import "fmt"

func getAverage(numbers ...float64) float64 {

	total := 0.0

	for _, value := range numbers {

		total += value
	}

	return total / float64(len(numbers))
}

func main() {

	fmt.Println("The average is:", getAverage(1, 2, 3))
}
