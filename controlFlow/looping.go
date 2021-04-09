// Exercise showing looping statements
// 1. The for statement
// 2. And more for statement... Um, yeah.
// 3. The "range" keyword with arrays/slices
//
// Notes:
//    Go only has the for loop
//    The ++ is not an expression, it is a statment

package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Nothing special about the statments between the ;'s
	for i, j := 0, 0; i < 10; i, j = i+1, j+2 { // Try replacing either increment assinment with the ++
		fmt.Println(i, j)
	}

	// The initial/conditional/incrementer are all optional... See where this is going?
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// Infinite for loop
	control := 0
	for {
		fmt.Println(control)
		if control >= 10 {
			break
		}
		control++
	}

	// And... The influence of Ken Thompson shows up. And he's such a smart guy too. Why? Oh, please, why?
BreakPoint:
	for m := 0; m < 5; m++ {
		for n := 0; n < 5; n++ {
			if m*n >= 7 {
				break BreakPoint // So we want to break BUT we'd only break the scope of the inner loop, hence the label
			}
		}
	}

	// And... Ken is back again! I knew he was smart. I knew I liked that guy.
	t := []int{1, 2, 3, 4, 5}
	for k, v := range t {
		fmt.Println(k, v)
	}
}
