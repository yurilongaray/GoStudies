package main

import (
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

func TestSum(t *testing.T) {

	const expected = 6
	received := Sum(2, 4)

	if received != expected {

		t.Errorf("Expected: %v; Received: %v", expected, received)
	}
}

func TestWrongSum(t *testing.T) {

	const expected = 6
	received := WrongSum(2, 4)

	if received != expected {

		t.Errorf("Expected: %v; Received: %v", expected, received)
	}
}
