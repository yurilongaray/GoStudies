package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	v, err := os.Open("names.txt")

	if err != nil {

		fmt.Println(err)
	}

	defer v.Close()

	sliceOfBytes, err2 := ioutil.ReadAll(v)

	if err2 != nil {

		fmt.Println(err2)
	}

	fmt.Println(string(sliceOfBytes))
}
