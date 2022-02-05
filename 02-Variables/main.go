package main

import "fmt"

func main() {
	// Variables declaration
	var number int

	// Variable definition
	number = 42

	// Variable declaration with initialization
	var name string = "John"

	// Dynamic variable declaration
	pi_number := 3.14

	// Mixed variable declaration
	var a, b, c = 3, 3.14, "Hello"

	fmt.Println(number, name)
	fmt.Printf("%T\n", pi_number)
	fmt.Println(a, b, c)

}
