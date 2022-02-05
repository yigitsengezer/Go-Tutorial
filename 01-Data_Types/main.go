package main

import "fmt"

func main() {

	// bool is a data type that represents a boolean value.
	var bool_var bool = true

	// string is a data type that represents a string of characters.
	var string_var string = "Hello, World!"

	// int is 32-bit or 64-bit integer depending on the architecture.
	var int_var int = 42
	// int8 is 8-bit integer(byte is alias for int8).
	var int8_var int8 = 127
	// int16 is 16-bit integer
	var int16_var int16 = 32767
	// int32 is 32-bit integer(rune is alias for int32).
	var int32_var int32 = 2147483647
	// int64 is 64-bit integer
	var int64_var int64 = 9223372036854775807

	// uint is 32-bit or 64-bit unsigned integer depending on the architecture.
	var uint_var uint = 42
	// uint8 is 8-bit unsigned integer
	var uint8_var uint8 = 255
	// uint16 is 16-bit unsigned integer
	var uint16_var uint16 = 65535
	// uint32 is 32-bit unsigned integer
	var uint32_var uint32 = 4294967295
	// uint64 is 64-bit unsigned integer
	var uint64_var uint64 = 18446744073709551615

	// float32 is 32-bit floating point number
	var float32_var float32 = 3.14
	// float64 is 64-bit floating point number
	var float64_var float64 = 3.141592653589793

	// complex64 is complex number with float32 real and imaginary parts
	var complex64_var complex64 = 3.14 + 2.72i
	// complex128 is complex number with float64 real and imaginary parts
	var complex128_var complex128 = 3.141592653589793 + 2.718281828459045i

	fmt.Println(bool_var)
	fmt.Println(string_var)
	fmt.Println(int_var, int8_var, int16_var, int32_var, int64_var)
	fmt.Println(uint_var, uint8_var, uint16_var, uint32_var, uint64_var)
	fmt.Println(float32_var, float64_var)
	fmt.Println(complex64_var, complex128_var)

}
