// Declare three variables that are initalized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initalize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c float64

	// Display the value of those variables.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Declare variables and initalize.
	// Using the short variable declaration operator.
	aa := 17
	bb := "test"
	cc := 12.34

	// Display the value of those variables.
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	// Perform a type conversion.
	ui := uint32(aa)

	// Display the value of that variable.
	fmt.Println(ui)

	var pi float32 = float32(3.14)
	fmt.Println(pi)
}
