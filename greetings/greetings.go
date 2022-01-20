package greetings

import (
	"fmt"
	"errors"
)

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. 
// This is known in Go as an exported name.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}


	// Go uses the value on the right to determine the variable's type
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}