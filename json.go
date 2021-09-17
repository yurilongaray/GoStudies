package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Values must have upper case on first letter to be seeing in a JSON
// ALWAYS REMMEBER THAT WHEN YOU WANT TO CHANGE A STRUCT TO JSON THE VALUES MUST HAVE THE FIRST LETTER UPPER CASED
// Ex.: Rename the name field to Name.
type Actor struct {
	Name  string
	Value float64
	Age   int
}

func main() {

	actor1 := Actor{
		Name:  "James Bond",
		Value: 20000.05,
		Age:   35,
	}

	value, err := json.Marshal(actor1)

	if err != nil {

		fmt.Println("Error", err)
	}

	// the value (return from json.Marshal) is in bits and must be stringified
	fmt.Println(value)
	// stringified bits
	fmt.Println(string(value))

	var actor2 Actor

	// transforming into a struct again to a new variable
	err1 := json.Unmarshal(value, &actor2)

	if err != nil {
		fmt.Println(err1)
	}

	fmt.Println(actor2)

	// os.Stdout is the terminal print
	encoder := json.NewEncoder(os.Stdout)

	// same as json.Marshal but the marshal returns to a variable and the encoder always go to that interface (os.Stdout)
	encoder.Encode(actor1)

	// stdout === Standard output (OS - Terminal)
	// stdin === Standard input (keyboard)
}
