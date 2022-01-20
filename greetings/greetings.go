package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. 
// This is known in Go as an exported name.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return name, errors.New("empty name")
	}

	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup, after global variables have been initialized. 
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
// randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}