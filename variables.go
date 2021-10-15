package main

import "fmt"

type hotdog int

// Test 1 https://www.youtube.com/watch?v=OINd35-02xo&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=22&ab_channel=AprendaGo
// Test 2 https://www.youtube.com/watch?v=9eMbGsKcWlc&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=37&ab_channel=AprendaGo
func main() {

	fmt.Println("test 1")

	x := 42
	var x1 int
	var x2 int = 42

	y := "James Bond"
	var y1 string = "James Bond"
	var y2 string

	z := true
	var z1 bool
	var z2 bool = true

	total := fmt.Sprintf("%v, %v, %v", x, y, z)
	var valueX hotdog

	fmt.Println("X: ", x, x1, x2)
	fmt.Println("Y: ", y, y1, y2)
	fmt.Println("Z: ", z, z1, z2)

	fmt.Println(total)
	fmt.Println(valueX)

	// show type
	fmt.Printf("%T, %T, %T, %T, %T", x, y, z, total, valueX)
	fmt.Println("")

	valueX = 99
	x = int(valueX)

	fmt.Println("X", x)
	fmt.Printf("%T", x)
	fmt.Println("")

	valueT := 100

	// show Binary, Decimal, Hexadecimal:
	fmt.Printf("%b, %d, %#x", valueT, valueT, valueT)
	fmt.Println("")

	const a1 = 99
	a2 := a1

	fmt.Printf("%v,%T\n", a1, a1)
	fmt.Printf("%v,%T\n", a2, a2)

	const a3 = 200
	// moving binary to sides (https://www.youtube.com/watch?v=USntMXkOihY&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=30&ab_channel=AprendaGo)
	const a4 = 200 << 1
	const a5 = 200 >> 1

	fmt.Printf("%b, %d, %#x\n", a3, a3, a3)
	fmt.Printf("%b, %d, %#x\n", a4, a4, a4)
	fmt.Printf("%b, %d, %#x\n", a5, a5, a5)

	// raw string literal
	const b2 = `
		A
		SINGLE
		TEST
	`

	fmt.Printf("%v\n", b2)

	// showing next years (https://www.youtube.com/watch?v=1IduyaGMO3g&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=30&ab_channel=AprendaGo)
	// more example https://www.youtube.com/watch?v=M07eDn7FyxI&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=36&ab_channel=AprendaGo
	const (
		nextYear1 = iota + 2021
		nextYear2
		_
		nextYear4
	)

	fmt.Printf("%v,%v,%v\n", nextYear1, nextYear2, nextYear4)

	var unknownVariable interface{}

	fmt.Printf("%v, %T\n", unknownVariable, unknownVariable)
}
