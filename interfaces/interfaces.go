// Exercise showing interfaces
// 1. Composition of interfaces
// 2. Type Conversion
// 3. The empty interface
// 4. Type switches
//
// Notes:
//   All the methods of an interface are referred to as that interface's "method set"
//      Method set of value is all methods with value receivers
//      Method set of pointer is all methods, regardless of receiver type
//   Variable declared as the concrete type of some interface, if declared by value,
//      receive all the methods of that interface that accept "value receivers". But
//      if that same variable is declared as a pointer to that concrete type, then
//      the variable receives the method set consisting of all methods that accept
//      pointer receivers and value receivers inclusive. Hence, variable declared as
//      pointers to an interface have a more flexible usage.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interfaces")

	// ConsoleWriter implicitly implements Writer
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Stuff to write"))

	var i interface{} = 0 // This is defined as an empty interface
	switch i.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("Dunno'")
	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Simple example of interface composition
type WriterCloser interface {
	Writer
	Closer
}

// Notice there is no explicit connection between the Writer and ConsoleWriter
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
