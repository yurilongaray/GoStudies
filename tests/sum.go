package main

func Sum(args ...int) int {

	total := 0

	for _, v := range args {

		total += v
	}

	return total
}

func WrongSum(args ...int) int {

	return 1 + 1
}
