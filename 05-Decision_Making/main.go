package main

func main() {
	var a int = 10

	// if statement syntax
	if a < 20 {
		println("a is less than 20")
	}

	// if else statement syntax
	if a > 20 {
		println("a is greater than 20")
	} else {
		println("a is less than 20")
	}

	// if else if statement syntax
	if a > 20 {
		println("a is greater than 20")
	} else if a < 20 {
		println("a is less than 20")
	} else {
		println("a is equal to 20")
	}

	// switch statement syntax
	switch a {
	case 10:
		println("a is equal to 10")
	case 20:
		println("a is equal to 20")
	default:
		println("a is not equal to 10 or 20")
	}

}
