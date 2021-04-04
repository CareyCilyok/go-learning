// Exercise showing typing and reflection
// 1. Show the "Pascal" variable declaration style
// 2. All variable MUST be used, hence unused varables create errors.
// 3. Runtime Type Information (RTTI) is available and can be accessed
// 4. Type inference via the ":=" assignment operator
// 5. Show the boolean type
// 6. Show usage of the "==" comparison operator
// 7. Show common arethmetic operations with typed output
// 8. While Go does type inference, it will not allow implicit type conversion

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var unused int	// Uncomment this line to show the "declared not used" error

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	// Introduction to the "strconv" package
	var j string = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

	// Type inference
	k := 42.5
	fmt.Printf("%v, %T\n", k, k)

	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	// The next line generates a "don't use Yoda conditions" warning... Don't you just love that?
	b1 := 1 == (1 - 0) // Try putting "1 == 1" and see what the Go warning is, hint: both sides of the comparison operator are identical
	b2 := 1 == 42
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", b2, b2)

	a := 1729
	b := 91
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // Will produce an integer result because both operands were inferred as ints
	fmt.Println(a % b)

	var c int = 10
	var d int8 = 3
	//fmt.Println(c + d) // Try uncommenting this line
	fmt.Println(c + int(d))
}
