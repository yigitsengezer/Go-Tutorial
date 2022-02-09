package main

import "fmt"

// Function as a return value:
func add(num1, num2 int) int {
	return num1 + num2
}

// Function without return value:
func print_add(num1, num2 int) {
	fmt.Println(num1 + num2)
}

// Function with multiple return values:
func find_max_and_min(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num1, num2
	}
	return num2, num1
}

// Function with call by reference:
func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

// A function can return another function:
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	x, y, radius float64
}

// Method is a function that is part of a struct:
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	// Call by value a function:
	sum := add(10, 20)
	fmt.Println(sum)

	// Call a function without return value:
	print_add(10, 20)

	// Call a function with multiple return values:
	max, min := find_max_and_min(10, 20)
	fmt.Println(max, min)

	// Call by reference a function:
	var num1, num2 int = 10, 20
	swap(&num1, &num2)

	// Call a function that returns another function:
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	// Call a method of a struct:
	c := Circle{x: 0, y: 0, radius: 5}
	area := c.getArea()
	fmt.Println(area)

}
