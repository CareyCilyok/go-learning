// Exercise showing Goroutines
// 1. Creation
// 2. Synchronization
// 3. Parallelism
//
// Notes:
//    Should prefer parmeter passing with Goroutines instead of closures
//    Like all other languges, mutexes aren't magic, use them properly
//    The "runtime" package has some very good interrogation and settings functionality
//    You shouldn't create goroutines in librarie, let your consumer control concurrency
//    Make sure you undertand your termination conditions for goroutines as they will create subtle resource drains
//    Use the compiler to check for race conditions
//       The "go run" command line has a "-race" flag which will help find them

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var counter = 0
var mutex = sync.RWMutex{}

func main() {
	fmt.Println("Goroutines")

	// Create a Goroutine ("green thread")
	go sayHello()                      // But the main is itself a Goroutine, hence if we don't wait a bit we'll never see this since the main will exit before this can execute
	time.Sleep(100 * time.Millisecond) // Comment out this line to see

	// Or how about...
	var msg = "Hola"
	go func() {
		fmt.Println(msg) // Go has closures, hence msg is closed over
	}()
	msg = "Hello" // BUT... Now we have the potential for a race
	time.Sleep(100 * time.Millisecond)

	// To avoid the race above by closing over the outer scoped msg, we just modify the lambda
	msg = "Hola"
	go func(msg string) {
		fmt.Println(msg) // We now pass in the msg instead of cloing over it
	}(msg)
	msg = "Hello" // No more race
	time.Sleep(100 * time.Millisecond)

	// But sleeps suck
	msg = "Wait on Hola"
	waitGroup.Add(1) // Says, add a goroutine to my waitGroup
	go func(msg string) {
		fmt.Println(msg)
		waitGroup.Done() // "Signal" that this goroutine is finished
	}(msg)
	msg = "Hello"    // No more race
	waitGroup.Wait() // Wait for the tracked goroutines to signal completion

	// Now, let's spin up several
	// We're just demonstrating how the mutex works but this is clearly a horrible way to do it
	//   -- All concurrency and parallelism is destroyed
	//   -- It's now a mess to read
	//   -- We're mixing synchronus calls to the locks with asynchronous calls to unlock
	numThreads := runtime.GOMAXPROCS(-1)
	fmt.Printf("Number of threads available are %v\n", numThreads)
	for i := 0; i < 10; i++ {
		waitGroup.Add(2) // Because we are about to add 2 goroutines
		mutex.RLock()
		go syncedHello()
		mutex.Lock()
		go increment()
	}
	waitGroup.Wait()
}

func sayHello() {
	fmt.Println("Howdy")
}

func syncedHello() {
	fmt.Printf("Howdy #%v\n", counter)
	mutex.RUnlock()
	waitGroup.Done()
}

func increment() {
	counter++
	mutex.Unlock()
	waitGroup.Done()
}
