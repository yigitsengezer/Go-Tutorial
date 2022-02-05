package main

import "fmt"

func main() {

	// Arithmetic operators
	// Add two operands
	add := 30 + 10
	// Subtract two operands
	sub := 30 - 10
	// Multiply two operands
	mul := 30 * 10
	// Divide two operands
	div := 30 / 10
	// Modulo two operands
	mod := 30 % 10
	fmt.Println(add, sub, mul, div, mod)

	// Relational operators
	// Check if two operands are equal
	is_equal := 10 == 10
	// Check if two operands are not equal
	is_not_equal := 10 != 10
	// Check if one operand is less than the other
	is_less_than := 10 < 10
	// Check if one operand is greater than the other
	is_greater_than := 10 > 10
	// Check if one operand is less than or equal to the other
	is_less_than_or_equal_to := 10 <= 10
	// Check if one operand is greater than or equal to the other
	is_greater_than_or_equal_to := 10 >= 10
	fmt.Println(is_equal, is_not_equal, is_less_than, is_greater_than, is_less_than_or_equal_to, is_greater_than_or_equal_to)

	// Logical operators
	// And Operator
	and_operator := true && false
	// Or Operator
	or_operator := true || false
	// Not Operator
	not_operator := !true
	fmt.Println(and_operator, or_operator, not_operator)

	// Bitwise operators
	// Bitwise AND
	bitwise_and := 5 & 3
	// Bitwise OR
	bitwise_or := 5 | 3
	// Bitwise XOR
	bitwise_xor := 5 ^ 3
	// Bitwise NOT
	bitwise_not := ^5
	// Bitwise Left Shift
	bitwise_left_shift := 5 << 2
	// Bitwise Right Shift
	bitwise_right_shift := 5 >> 2
	fmt.Println(bitwise_and, bitwise_or, bitwise_xor, bitwise_not, bitwise_left_shift, bitwise_right_shift)

	// Assignment operators
	// Simple assignment
	simple_assignment := 10
	// Addition assignment
	addition_assignment := 10
	addition_assignment += 10
	// Subtraction assignment
	subtraction_assignment := 10
	subtraction_assignment -= 10
	// Multiplication assignment
	multiplication_assignment := 10
	multiplication_assignment *= 10
	// Division assignment
	division_assignment := 10
	division_assignment /= 10
	// Modulo assignment
	modulo_assignment := 10
	modulo_assignment %= 10
	// Bitwise AND assignment
	bitwise_and_assignment := 10
	bitwise_and_assignment &= 10
	// Bitwise OR assignment
	bitwise_or_assignment := 10
	bitwise_or_assignment |= 10
	// Bitwise XOR assignment
	bitwise_xor_assignment := 10
	bitwise_xor_assignment ^= 10
	// Bitwise Left Shift assignment
	bitwise_left_shift_assignment := 10
	bitwise_left_shift_assignment <<= 10
	// Bitwise Right Shift assignment
	bitwise_right_shift_assignment := 10
	bitwise_right_shift_assignment >>= 10
	fmt.Println(simple_assignment, addition_assignment, subtraction_assignment, multiplication_assignment, division_assignment, modulo_assignment, bitwise_and_assignment, bitwise_or_assignment, bitwise_xor_assignment, bitwise_left_shift_assignment, bitwise_right_shift_assignment)

	// Miscellaneous operators
	// Adress of operator
	address_of_operator := &simple_assignment
	fmt.Println(address_of_operator)
	// Dereference operator
	dereference_operator := *address_of_operator
	fmt.Println(dereference_operator)
}
