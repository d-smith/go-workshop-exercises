// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/b6-FNFOToO

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import (
	"fmt"
)

// Declare a type named user.
type user struct {
	name  string
	email string
}

// Create a function that changes the value of one of the user fields.
func funcName(u *user, newEmail string) {
	// Use the pointer to change the value that the
	// pointer points to.
	u.email = newEmail
}

// main is the entry point for the application.
func main() {
	// Create a variable of type user and initialize each field.
	u := user{"foo", "foo@foo.com"}

	// Display the value of the variable.
	fmt.Printf("%v\n", u)

	// Share the variable with the function you declared above.
	funcName(&u, "new@mail.com")

	// Display the value of the variable.
	fmt.Printf("%v\n", u)
}
