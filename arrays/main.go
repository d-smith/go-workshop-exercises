// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/ggjjRPzhAB

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
package main

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare an array of 5 strings set to its zero value.
	var stringArr [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	stringArr2 := [5]string{"a", "b", "C", "d", "E"}

	// Assign the populated array to the array of zero values.
	stringArr = stringArr2

	// Iterate over the first array declared.
	for _, v := range stringArr {
		println(v)
	}
}
