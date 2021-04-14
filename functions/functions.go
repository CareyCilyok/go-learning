// Exercise showing functions
// 1. Basic syntax
// 2. Parameters
// 3. Return values
// 4. Anonymous functions
// 5. Functions as type (first class functions)
// 6. Methods
//
// Notes:
//   func name(<params>) (<name> <return type>) { <body> }
//   Go supports variatic parameters with the "..." syntax, similar to varargs
//   Variatic parameters must be the last parameter
//   Go will allow the access of a local value to outer scope when the stack for that func is destroyed
//   You can also name return values
//   We can also have multiple return values

package main

import (
	"fmt"
	"math"
)

// Define a type and a "method" on that type
type GeoLoc struct {
	lat float64
	lon float64
}

func (location GeoLoc) print() {
	fmt.Println(location.lat, location.lon)
}

func main() {
	fmt.Println("Functions")

	// Can pass any number of typed parameters
	sayMessage("Just gonna print this")

	// Simple return value
	fmt.Println(add(1, 2))

	// Simple example using variatic paraeters
	fmt.Println(addMany(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// Read the function below - in Go this is totally safe but freaky
	g := totallyMadeupExampleOfSomethingThatShouldFreakoutMostDevelopers(10)
	fmt.Println(*g)

	fmt.Println(namedReturnValue(1729))

	result, err := divide(1.0, 0.0) // Is going to try to return 1.0 / 0.0
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Here is an anonymous function (a function as a type)
	func() { fmt.Println("Lambda") }() // Declare and execute via the ending ()

	// Example of a method
	loc := GeoLoc{
		lat: 0.0,
		lon: 0.0,
	}
	loc.print()
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

// Notice here that since x and y are both int, we can simply separate them with comma and then add tthe type
func add(x, y int) int {
	return x + y
}

// Notice the use of the "..." and how to access them parameters
func addMany(numbers ...int) int {
	fmt.Println(numbers)
	result := 0

	for _, num := range numbers {
		result += num
	}

	return result
}

// Note that the return type is a pointer to an int
func totallyMadeupExampleOfSomethingThatShouldFreakoutMostDevelopers(arg int) *int {
	result := arg + arg // We really don't care
	return &result      // THIS is the interesting part - result is local
}

// Notice we name our return value in the func declaration and then use it
func namedReturnValue(x float64) (result float64) {
	result = math.Pow(x, 2) // Now just use "result"
	return                  // No need to specify what we are returning, as we named in in the func decl
}

// Notice the multiple return types
// Returning error status this way is a very idomatic way to handle this in Go
func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return x / y, nil
}
