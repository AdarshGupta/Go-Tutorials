package greetings

import "fmt"

// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. 
// This is known in Go as an exported name.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}