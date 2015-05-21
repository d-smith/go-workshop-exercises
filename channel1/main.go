// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/G7O-DnJrEA

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

// Add imports.
import (
	"fmt"
	"sync"
)

// Declare a wait group variable.
var wg sync.WaitGroup

// Declare a function for the goroutine. Pass in a name for the
// routine and a channel of type int used to share the value back and forth.
func goroutine(name string, c chan int) {
	for {
		// Receive on the channel, waiting for the integer.
		n, ok := <-c

		// Check if the channel is closed.
		if !ok {
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("%s is done\n", name)
			return
		}

		// Display the integer value received.
		fmt.Printf("%s recieved value %d\n", name, n)
		// Check is the value from the channel is 10.
		if n == 10 {
			// Close the channel.
			close(c)
			// Call done on the waitgroup.
			wg.Done()
			// Display the goroutine is finished and return.
			fmt.Printf("%s has closed the channel and is finished\n", name)
			return
		}

		// Increment the value by one and send in back through
		// the channel.
		n++
		c <- n
	}
}

// main is the entry point for all Go programs.
func main() {
	// Declare and initialize an unbuffered channel
	// of type int.
	mychan := make(chan int)

	// Increment the wait group for each goroutine
	// to be created.
	wg.Add(2)

	// Create the two goroutines.
	go goroutine("thing1", mychan)
	go goroutine("thing2", mychan)

	// Send the initial integer value into the channel.
	mychan <- 0

	// Wait for all the goroutines to finish.
	wg.Wait()
}
