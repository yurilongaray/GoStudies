package main

import (
	"fmt"
	"runtime"
)

// run using: go run *.go
func main() {

	fmt.Println("This is from main")
	fmt.Println(runtime.NumCPU())
	fmt.Println("You OS is: ", runtime.GOOS)
	fmt.Println("You ARCH is: ", runtime.GOARCH)
	Other()
}
