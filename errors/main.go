// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rt3O-7ndtJ

// Create two error variables, one called ErrInvalidValue and the other
// called ErrAmountTooLarge. Provide the static message for each variable.
// Then write a function called checkAmount that accepts a float64 type value
// and returns an error value. Check the value for zero and if it is, return
// the ErrInvalidValue. Check the value for greater than $1,000 and if it is,
// reutrn the ErrAmountTooLarge. Write a main function to call the checkAmount
// function and check the return error value. Display a proper message to the screen.
package main

// Add imports.
import (
	"errors"
	"fmt"
)

// Declare an error variable named ErrInvalidValue using the New
// function from the errors package.
var ErrInvalidValue = errors.New("Invalid value")

// Declare an error variable named ErrAmountTooLarge using the New
// function from the errors package.
var ErrAmountTooLarge = errors.New("Too large of an amount")

// Declare a function named checkAmount that accepts a value of
// type float64 and returns an error interface value.
func checkAmount(amount float64) error {
	// Is the parameter equal to zero. If so then return
	// the error variable.
	if amount == 0 {
		return ErrInvalidValue
	}

	// Is the parameter greater than 1000. If so then return
	// the other error variable.
	if amount > 1000 {
		return ErrAmountTooLarge
	}

	// Return nil for the error value.
	return nil
}

// main is the entry point for the application.
func main() {
	// Call the checkAmount function and check the error. Then
	// use a switch/case to compare the error with each variable.
	// Add a default case. Return if there is an error.
	amount := 2.0

	err := checkAmount(amount)
	switch err {
	case ErrAmountTooLarge:
		fmt.Println(ErrAmountTooLarge.Error())
		return
	case ErrInvalidValue:
		fmt.Println(ErrInvalidValue.Error())
		return

	}

	// Display everything is good.
	fmt.Println("All is in order")
}
