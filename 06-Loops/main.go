package main

import "fmt"

func main() {

	// For loop syntax
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Go do not have while loop we have to use for loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Break statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	// Continue statement
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	// Goto statement
	i = 0
	for {
		if i == 5 {
			goto end
		}
		fmt.Println(i)
		i++
	}
end:
	fmt.Println("End")

}
