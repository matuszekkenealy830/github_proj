package main

import "fmt"

func main() {
	// Declare a variable to hold the user's name
	var name string

	// Prompt the user to enter their name
	fmt.Println("What is your name?")

	// Read the user's input and store it in the name variable
	fmt.Scan(&name)

	// Print a message to the user with their name
	fmt.Println("Hi", name, "!")
}
