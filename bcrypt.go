package main

// go get golang.org/x/crypto/bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "CorrectPassword"

	// sb is a commum name for slice of bytes
	// GenerateFromPassword(the value to hash, the complexity by the computational cost)
	sb, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(password)))            // nil because its ok the hash and the password
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte("IncorrectPassword"))) // error message because its not equal
}
