// Exercise showing maps and structs
// 1. Creating maps (reference types)
// 2. Manipulating maps
// 3. Introduction to the "ok" assingment syntax
// 4. Creating structs (value types)
// 5. Anonymous structs
// 6. Struct embedding (in lieu of implementation inheritance, hence composition)
// 7. Struct tags

package main

import (
	"fmt"
	"reflect"
)

// Declare a fully exportable struct.
// Remember upper case names are exported while lower case name are not
type PoolMatch struct {
	Player1  string
	Player2  string
	GameType string
}

// Next two structs are to show embedding (composition)
// Notice the "tag" on the name, example of using the tag later down
type Animal struct {
	Name   string `required max: "100"`
	Origin string
}

// While we don't have "implementation inheritance" we can use composition - notice the embedded Animal struct
// We'd say here that Bird "has a" Animal but we can't say a Bird "is a" Animal yet - We're going to need formal interfaces (later)
type Bird struct {
	Animal
	CanFly bool
}

func main() {
	// Literal definition (don't forget we could use "make(map[string]int) but it's less common)
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
	fmt.Println(digitNames)

	// Or get a digit value by name
	fmt.Println(digitNames["Seven"])

	// But what if we need to interrogate if a value exists in the map
	_, ok := digitNames["i-Squared"] // Is this in the map? (notice the use of the placeholder syntax as well)
	fmt.Println(ok)                  // The "ok" is a boolean letting us know if the value was in this map

	// Name field assignment syntax used. You can use purely positional syntax without names but it's not recommended (think about later changes to the struct)
	aPoolMatch := PoolMatch{
		Player1:  "Effern Reyes",
		Player2:  "Shane Van Boening",
		GameType: "One Pocket",
	}
	fmt.Println(aPoolMatch)

	// Standard sort of struct accessor syntax
	fmt.Println(aPoolMatch.GameType)

	// Also, you can have an anonymous structs - rare but maybe we would use these for things like data projections
	aFullName := struct {
		first  string
		middle string
		last   string
	}{first: "John", middle: "Quincy", last: "Adams"}
	fmt.Println(aFullName)

	// Use reflection to access the tag we applied in the definition
	animalType := reflect.TypeOf(Animal{})
	field, _ := animalType.FieldByName("Name")
	fmt.Println(field.Tag)
}
