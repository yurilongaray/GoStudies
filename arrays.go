package main

import "fmt"

func main() {

	// This is array (preferable to use slice). In Go, arrays are not used as much as slices.
	// Arrays have fixed length, slice is dynamic
	var x [5]int

	fmt.Printf("%v, %T\n", x, x)

	// This is slice, slices are made of arrays
	y := []int{1, 2, 3, 4, 5}

	fmt.Printf("%v, %T\n", y, y)

	// same as push from javascript
	y = append(y, 99, 102, 300)
	total := 0

	// _ ignores the value
	for _, value := range y {

		total += value
		fmt.Println(value)
	}

	fmt.Println("Total", total)

	// slicing a slice
	slicedValues := y[2 : len(y)-1]

	fmt.Println(slicedValues)

	// without using range
	for i := 0; i < len(y); i++ {

		fmt.Println(y[i])
	}

	// removing a value from a slice
	newSlice := append(y[:1], y[3:]...)

	fmt.Println(newSlice)

	// appending two slices
	arrayResult := append(y, slicedValues...)

	fmt.Println(arrayResult)

	sliceByMake := make([]int, 5, 10)

	fmt.Println("sliceByMake", sliceByMake)
	fmt.Println("Cap", cap(sliceByMake))

	ss := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7},
	}

	fmt.Println(ss)
	fmt.Println("ss[2][1]", ss[1][2])

	testMap := map[string]int{
		"test1":  22,
		"test2":  33,
		"test33": 112,
	}

	fmt.Println(testMap)
	fmt.Println(testMap["test2"])

	testMap["test3"] = 99

	fmt.Println(testMap)

	// the second prop below shows that the value doesnt exist in map (COMMA OK IDIOM)
	prop, ok := testMap["test99"]
	fmt.Println(prop, ok)

	if propEx, existsEx := testMap["test99"]; !existsEx {

		fmt.Println("Doesnt Exists", propEx)
	} else {

		fmt.Println("Exists", propEx)
	}

	// same as Object.entries from javascript
	for key, value := range testMap {

		fmt.Println("key, value", key, value)
	}

	// deleting value from map
	delete(testMap, "test33")

	fmt.Println(testMap)

	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	newB := append(b, 52, 53, 54, 55)

	newB = append(b, newB...)

	fmt.Println(newB)

	// changing the base array from slices (both slices are pointed to the same local in memory)
	slice1 := make([]int, 10, 20)
	slice2 := append(slice1, 1, 2, 3)

	fmt.Println(slice1, slice2)

	// OR slice1[1] = 99
	slice2[1] = 99

	fmt.Println(slice1, slice2)
}
