// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/-JBSUoux-v

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
	m := make(map[string]int)

	// Intialize some data into the map.
	m["a"] = 1
	m["b"] = 2

	// Display each key/value pair.
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
