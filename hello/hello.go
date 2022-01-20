// In Go, code executed as an application must be in a main package.
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	message, err := greetings.Hello("")

	// If an error was returned, print it to the console and
    // exit the program.
	if err != nil {
		// log package's Fatal function to print the error and stop the program.
		log.Fatal(err)
	}
	fmt.Println(message)
}