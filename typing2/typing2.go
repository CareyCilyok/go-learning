// Exercise showing more primative types and operations
// 1. Show logical and shift bit operations
// 2. Show various floating types
// 3. Show type mismatches for floating types
// 4. Show complex types
// 5. Show the string types
// 6. Show the immutability of strings
// 7. Strings as a collection of bytes or the Go "slice"
// 8. Introduction of "runes"

package main

import (
	"fmt"
	"math"
)

func main() {
	// Logical bit operations
	a := 8              // 1000
	b := 3              // 0011
	fmt.Println(a & b)  // 0000 = 0
	fmt.Println(a | b)  // 1011 = 11
	fmt.Println(a ^ b)  // 1011 = 11
	fmt.Println(a &^ b) // 1000 = 8

	// Bit shifting operations
	fmt.Println(a << 3) // 64
	fmt.Println(a >> 3) // 1

	var c float32 = 2.99792458e8             // Yeah, that "c"
	var m float64 = 5.9722e24                // Rest mass of the Earth (within reasonable error) - Yes, I know it will still fit in a float32 but that doesn't help with the example
	fmt.Println(m * math.Pow(float64(c), 2)) // E = mc^2 - but remove the float64 conversion for "c" and it won't compile

	// Complex numbers are first class types
	var complexNumber complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", complexNumber, complexNumber)
	fmt.Printf("%v, %T\n", real(complexNumber), real(complexNumber))
	fmt.Printf("%v, %T\n", imag(complexNumber), imag(complexNumber))

	// Strings and text
	quote := "I'd have put and 'e' on creat. - Ken Thompson"
	//quote[29] = "e" // Trying uncommenting this line to get the error about the immutability of strings
	fmt.Printf("%v, %T\n", quote, quote)

	// String can also just be treated as arrays
	fmt.Printf("%v, %T\n", string(quote[18]), quote[18])

	// Examine the string as a "byte slice"
	quoteArray := []byte(quote)
	fmt.Printf("%v, %T\n", quoteArray, quoteArray)

	// Runes are UTF-32 not UTF-8
	var r rune = 'p'
	fmt.Printf("%v, %T\n", r, r)
}
