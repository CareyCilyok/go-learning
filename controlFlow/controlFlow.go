// Exercise showing control statements
// 1. If statements
// 2. Usage with the "ok" result
// 3. Switch statements
// 4. A "type switch"
//
// Notes:
//   Go implements short circuiting as expected and implents lazy evaluation
//   The assumed else and else if syntax is also available
//   The switch/case formatting is really the first time I don't like fmt.  :-(
//   Switch statements have an implicit "break", hence Go has an explicit "fallthrough" keyword to force that behavior
//   The "break" keyword is actually used but its typical usage is for an early break out of a case condition

package main

import (
	"fmt"
	"math"
)

func main() {
	if true { // Clearly the most simple example
		fmt.Println("true")
	}

	digitNames := map[string]int{
		"Zero":  0,
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
		"Eight": 8,
		"Nine":  9,
	}

	// A very common usage though is to combine an assigment separated from the conditional by ;
	if digit, ok := digitNames["Zero"]; ok {
		fmt.Println(digit) // digit is only scope within this block
	} else {
		fmt.Println("Nope")
	}

	// Switch is also available and works as we'd expect
	// Notice that we can have an optional initializer followed by the test value
	switch taxicab := math.Pow(9, 3) + math.Pow(10, 3); taxicab {
	case 1: // Can match against a single value
		fmt.Println("Not the first taxi cab number")
	case 31, 57: // Can test against a list of matches
		fmt.Println("Not the first taxi cab number")
	case 1729:
		fmt.Println("The first taxi cab number")
	default:
		fmt.Println("Huh?")
	}

	// switch also works "without" at test value...
	taxicab := math.Pow(9, 3) + math.Pow(10, 3)
	switch {
	case taxicab < 1729:
		fmt.Println("Nope")
	case taxicab > 1729:
		fmt.Println("Nope")
	default:
		fmt.Println("Must be the first Taxicab number")
	}

	// This use of a switch also introduces the "interface" and built-in reflection
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("A type we didn't check for")
	}
}
