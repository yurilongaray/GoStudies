package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	logFile, _ := os.Create("log.txt")

	defer logFile.Close()

	v, err := os.Open("NON_EXISTENT.txt")

	defer v.Close()

	if err != nil {

		// with this bellow, all the log (log.*) are saved in a file
		log.SetOutput(logFile)

		// always choose a log to save your erros
		log.Println("log.Println: ", err)
		log.Panicln(err)
		log.Fatalln("log.Fatalln: ", err)

		fmt.Println("fmt.Println: ", err)

		panic(err)
	}
}

/* Creating a custom error */
// - 1. errors.New
// - 2. var errors.New
// - 3. fmt.Errorf
// - 4. var fmt.Errorf
// - 5. type + method = error interface
