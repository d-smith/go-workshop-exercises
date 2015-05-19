// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initalize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import (
	"fmt"
)

// Add user type and provide comment.
type User struct {
	name  string
	email string
	age   int
}

// main is the entry point for the application.
func main() {
	// Declare variable of type user and init using a composite literal.
	user1 := User{"foo", "foo@foo.com", 33}

	// Display the field values.
	fmt.Println(user1.age)
	fmt.Println(user1.email)
	fmt.Println(user1.age)

	// Declare a variable using an anonymous struct.
	var foo = struct {
		x int
		y int
	}{1, 2}

	// Display the field values.
	fmt.Println(foo.x)
	fmt.Println(foo.y)
}
