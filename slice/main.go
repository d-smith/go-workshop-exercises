// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/mPKmyGNR2L

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import (
	"fmt"
)

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var ints []int

	// Appends numbers to the slice.
	for i := 0; i < 10; i++ {
		ints = append(ints, i)
	}

	// Display each value in the slice.
	for _, v := range ints {
		fmt.Printf("%d\n", v)
	}

	// Declare a slice of strings and populate the slice with names.
	strings := []string{"a", "b", "c"}

	// Display each index position and slice value.
	for i, v := range strings {
		fmt.Printf("idx %d is %s\n", i, v)
	}

	// // Take a slice of index 1 and 2 of the slice of strings.
	s2 := strings[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range s2 {
		fmt.Printf("idx %d is %s\n", i, v)
	}
}
