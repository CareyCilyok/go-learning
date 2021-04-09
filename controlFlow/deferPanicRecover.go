// Exercise showing defer, panic and recovery
// 1. The defer keyword
// 2. The panic keyword
// 3. The recover function
// 4. Introduce an anonymous function (lambda) to handle the recovery from a panic, see the recoveryExample
//
// Notes:
//    defer defers the statment AFTER the end of the function but before the function returns
//    defer statements run in stack (LIFO) order
//    panic statements happen AFTER deferred statements, hence defers that close resources will execute even under a panic
//

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// simpleDefer()
	// webRequestExample()
	// deferAssigenmentTime()

	// Commented out for obvious reasons but you can uncomment it just to see
	// panicExample()

	// When you run this example, just open a browser and look at address "localhost:8080"
	// webServer()

	// We call a panic in this example but recover from it
	recoveryExample()
	fmt.Println("After")
}

func simpleDefer() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}

// Taken directly from the Go documentation
func webRequestExample() {
	res, err := http.Get("http://www.google.com/robots.txt")
	defer res.Body.Close() // Now we defer the close because we don't want to forget
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body) // We're using the resource again
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
}

func deferAssigenmentTime() {
	output := "A. Will we print this?"
	fmt.Println(output)
	output = "B. Or... will we print this?"

	// We will print output A because defer makes all assignment with the values at that time
}

func panicExample() {
	// Will generate a stack trace and all that sort of panicky shit
	panic("Catastrophe")
}

func webServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Look! A web server!"))
	})

	// Change the port to a blocked port to see a practical use of panic
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func recoveryExample() {
	fmt.Println("Funtion entry")

	// Now, defer a lambda that "recovers"
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			// But if we look at err and decide we can't deal with it, we could call panic again
		}
	}()

	panic("Some reason we decided to pull the ripcord")
	fmt.Println("Function exit") // We won't see this
}
