// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/p0vlsW5sVL

// Declare a struct type to maintain information about a user. Declare a function
// that creates value of and returns pointers of this type and an error value. Call
// this function from main and display the value.
//
// Make a second call to your function but this time ignore the value and just test
// the error value.
package main

// Add imports.
import "fmt"

// Declare a type named user.
type user struct {
	name  string
	email string
}

// Declare a function that creates user type values and returns a pointer
// to that value and an error value of nil.
func NewUser(name string, email string) (*user, error) {
	if name == "NotAName" {
		return nil, fmt.Errorf("Error creating user")
	}
	return &user{name, email}, nil
}

// main is the entry point for the application.
func main() {
	// Use the function to create a value of type user. Check
	// the error being returned.
	user, err := NewUser("foo", "foo@foo.com")

	// Display the value that the pointer points to.
	fmt.Printf("%v\n", *user)

	// Call the function again and just check the error.
	_, err = NewUser("NotAName", "")
	if err != nil {
		println(err.Error())
	}

}
