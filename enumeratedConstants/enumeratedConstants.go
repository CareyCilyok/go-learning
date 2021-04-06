// Exercise showing enumerated constants and related feature
// 1. Show an enumerated constant
// 2. The use of the iota Go function
// 3. The use of the Go "_" write only target
// 4. The use of operations to create new constant values
// 5. Implicitly show that Go repeats the pattern it sees in the constant assignments

package main

import "fmt"

// Example taken from "Effective Go" on golang.org
const (
	_  = iota // Ignore the zero value of the iota function to the Go write only target
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// You can easily see from the enumerated constant above that creating bitwise values would be trivial

func main() {
	fileSize := 4000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)
}
