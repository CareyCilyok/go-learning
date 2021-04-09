// Exercise showing pointers
// 1. The & operator
// 2. The * operator
// 3. The new() function
// 4. The nil pointer
// 5. Precedence of the (), * and . operators
// 6. A little syntactic sugar from the Go compiler
//
// Notes:
//   Not showing any example but pointer arithmatic is NOT allowed in Go
//   You can totally fuck up your life and go get the https://golang.org/pkg/unsafe package but let's be honest, then just go use C

package main

import (
	"fmt"
)

func main() {
	var v1 int = 1279
	var v2 int = v1              // Both v1 and v2 are value types
	var vp *int = &v1            // * here means declare a pointer and & is the "address of" operator
	fmt.Println(v1, v2, vp, *vp) // * here is the derference operator
	v1 = 5150                    // Now set the original value to something new
	fmt.Println(v1, v2, vp, *vp) // What do we get now?

	var s *exStruct = &exStruct{f: 5150} // Declare a pointer to a struct
	fmt.Println(s)                       // This is instructive - will say "this is address of an object that holds a structure with a field that has the value 5150"

	// We can use "new" as well
	var s2 *exStruct     // Lets declare a pointer
	fmt.Println(s2)      // The uninitialize value is instructive, "nil" (Someone on the Go team liked Pascal, jus' sayin')
	s2 = new(exStruct)   // Mmm... Heap allocation
	fmt.Println(s2)      // Again, the output is instructive
	(*s2).f = 8675309    // "Jenny, Jenny, who can I turn to?"
	fmt.Println((*s2).f) // In Go the * operator is lower precedence than the . operator, hence the parens

	// But that was awkward, so the Go designers help us out - C/C++ programmers are about to shit themselves
	s2.f = 911
	fmt.Println(s2.f)
}

type exStruct struct {
	f int
}
