package main

import (
	"fmt"
	"sort"
)

// https://pkg.go.dev/sort

type HumamBeing struct {
	Name   string
	Age    int
	Salary float64
}

type byName []HumamBeing

func (x byName) Len() int           { return len(x) }
func (x byName) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x byName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byAge []HumamBeing

func (x byAge) Len() int           { return len(x) }
func (x byAge) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x byAge) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type bySalary []HumamBeing

func (x bySalary) Len() int           { return len(x) }
func (x bySalary) Less(i, j int) bool { return x[i].Salary < x[j].Salary }
func (x bySalary) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	sliceOfStrings := []string{"jump", "run", "walk", "run", "await"}

	sort.Strings(sliceOfStrings)

	fmt.Println(sliceOfStrings)

	sliceOfInts := []int{1, 4, 2, 55, 42}

	sort.Ints(sliceOfInts)

	fmt.Println(sliceOfInts)

	humans := []HumamBeing{
		{
			Name:   "Maria",
			Age:    66,
			Salary: 3000.29,
		},
		{
			Name:   "Ana",
			Age:    21,
			Salary: 9000.20,
		},
		{
			Name:   "Pedro",
			Age:    18,
			Salary: 1500.29,
		},
	}

	fmt.Println("Values to be sorted", humans)

	// sorting slice of structs
	sort.Slice(humans, func(i, j int) bool {

		return humans[i].Age < humans[j].Age
	})

	fmt.Println("sorted by sort.Slice (an easy way to sort)", humans)

	// creating a sort
	sort.Sort(byName(humans))

	fmt.Println("sorted by sort.Sort by NAME", humans)

	// creating another sort
	sort.Sort(byAge(humans))

	fmt.Println("sorted by sort.Sort by AGE", humans)

	// creating another sort
	sort.Sort(bySalary(humans))

	fmt.Println("sorted by sort.Sort by SALARY", humans)
}
