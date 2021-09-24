package main

import (
	"fmt"
	"testing"
)

type test struct {
	data   []int
	answer int
}

/* Testing multiple values */
func TestSumWithTable(t *testing.T) {

	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{1, 2, 3}, 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}

	for _, row := range tests {

		expected := row.answer
		received := Sum(row.data...)

		if received != expected {

			t.Errorf("Expected: %v; Received: %v", expected, received)
		}
	}
}

/* Tests As Example (they are like a test and they are shown into GoDocs, but the package must be main) - https://go.dev/blog/examples*/
func ExampleSum() {

	fmt.Println(Sum(9, 9))
	// Output: 18
}

func TestSum(t *testing.T) {

	const expected = 6
	received := Sum(2, 4)

	if received != expected {

		t.Errorf("Expected: %v; Received: %v", expected, received)
	}
}

/*
func TestWrongSum(t *testing.T) {

	const expected = 6
	received := WrongSum(2, 4)

	if received != expected {

		t.Errorf("Expected: %v; Received: %v", expected, received)
	}
}
*/

/* Run using "go test -bench Sum" or "go test -bench ." to run all */
func BenchmarkSum(b *testing.B) {

	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{1, 2, 3}, 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}

	for i := 0; i < b.N; i++ {

		for _, v := range tests {

			Sum(v.data...)
		}
	}
}
