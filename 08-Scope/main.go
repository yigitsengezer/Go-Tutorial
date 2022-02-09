package main

import "fmt"

// Global variables can be accessed by any function in the program.
var g int

func main() {
	// Local variables only exist within the function they are declared in.
	var l int = 10
	fmt.Println(l)

	g = 40
	fmt.Println(g)

	// If local and global variables have the same name, the local variable has priority.
	var g int = 50
	fmt.Println(g)

}
