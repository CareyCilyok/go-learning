// Exercise showing arrays and slices and related features
// 1. Creation
// 2. Built-in functions
// 3. Arrays are value types
// 4. Arrays are fixed size at compile
// 5. Slices are like arrays but are not fixed size at compile time
// 6. Slices are by default reference types
// 7. Different ways to create slices
// 8. Introduction to the "make" function to create a slice
// 9. Introduction of the "append" function to grow a slice
// 10. Manipulating slices - including a gotcha with slices manipulating them to remove a value in the middle
// 11. Another use of the "..." to implement a "spread", will decompose a slice into individual components

package main

import "fmt"

func main() {
	temps := [...]int{85, 86, 90} // Use of the ... operator allows for the size initialization to be inferred
	fmt.Printf("Temeratures: %v, Count = %v\n", temps, len(temps))

	var humidities [3]float32 // Here the size is explicit
	humidities[0] = 80.5
	humidities[1] = 70.2
	humidities[2] = 96.5
	fmt.Printf("Humitities: %v, Count = %v\n", humidities, len(humidities))

	// Arrays are VLUE types in Go...
	a := [...]int{1, 2, 3} // Again, inferred size at compile time
	b := a                 // Not a reference, makes an entire new copy of the data in a
	b[1] = 5               // Hence, this changes b[1] but not a[1]
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	// But Go has a pointer operator - "&"
	c := &a  // The creates a reference to a
	c[1] = 5 // Hence, this changes c[1] and a[1]
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", c)

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}                                       // It's subtle but this is a slice due to lack of a predefined size
	fmt.Printf("%v, Length of s1 = %v, Capacity of s1 = %v\n", s1, len(s1), cap(s1)) // Slices have a length and an underlying capacity of the backing array

	// More ways to create slices which is illustrative of the name
	s2 := s1[:]   // Same as s1
	s3 := s1[4:]  // s1 from 5th element to the end
	s4 := s1[:4]  // s1 first 4 elements
	s5 := s1[2:7] // s1 the 3rd, 4th, 5th, 6th and 7th elements

	// This output will show how the slices were assigned and the varying lengths and capacities
	fmt.Printf("%v, Length of s2 = %v, Capacity of s2 = %v\n", s2, len(s2), cap(s2))
	fmt.Printf("%v, Length of s3 = %v, Capacity of s3 = %v\n", s3, len(s3), cap(s3))
	fmt.Printf("%v, Length of s4 = %v, Capacity of s4 = %v\n", s4, len(s4), cap(s4))
	fmt.Printf("%v, Length of s5 = %v, Capacity of s5 = %v\n", s5, len(s5), cap(s5))

	// By default, slices are reference types so...
	s1[4] = 23
	fmt.Printf("%v, Length of s2 = %v, Capacity of s2 = %v\n", s2, len(s2), cap(s2))
	fmt.Printf("%v, Length of s3 = %v, Capacity of s3 = %v\n", s3, len(s3), cap(s3))
	fmt.Printf("%v, Length of s4 = %v, Capacity of s4 = %v\n", s4, len(s4), cap(s4))
	fmt.Printf("%v, Length of s5 = %v, Capacity of s5 = %v\n", s5, len(s5), cap(s5))

	s6 := make([]int, 3, 4) // If we leave off the 4, then we get a capacity equal to the length
	fmt.Printf("%v, Length of s6 = %v, Capacity of s6 = %v\n", s6, len(s6), cap(s6))

	s6 = append(s6, 1, 3, 5, 7, 11) // Variatic function allows you to append as many values as you like
	fmt.Printf("%v, Length of s6 = %v, Capacity of s6 = %v\n", s6, len(s6), cap(s6))

	// Using the slice operations we can remove from the beginning, the end and the middle
	// But you have to be careful...
	s7 := []int{1, 2, 3, 4, 5}
	fmt.Println(s7)
	s8 := s7[1:] // Trim from the beginning
	fmt.Println(s8)
	s9 := s7[:len(s7)-1] // Trim from the end (they're 0 based)
	fmt.Println(s9)
	s10 := append(s7[:2], s7[3:]...) // Remove the middle element - try removing the spread operator and see the compiler error
	fmt.Println(s10)

	// But slices are reference types so, guess what happened to s7
	fmt.Println(s7)
}
