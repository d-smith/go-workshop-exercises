// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/OkIHsVwMQ7

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// Marshal the slice into pretty print strings and display each element.
package main

// Add imports.
import (
	"encoding/json"
	"fmt"
	"os"
)

// Declare a struct type named user with two fields. Name of type string and
// Email of type string. Add tags for each field for the unmarshal call.
type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// main is the entry point for the application.
func main() {
	// Use the os package to Open the JSON file. Check for errors.
	f, err := os.Open("input.json")
	if err != nil {
		fmt.Printf("error opening input file: %s\n", err.Error())
		return
	}

	// Schedule the file to be closed once the function returns.
	defer f.Close()

	// Declare a nil slice of user struct types.
	var users []user
	// Decode the JSON from the file into the slice. Check for errors.
	err = json.NewDecoder(f).Decode(&users)
	if err != nil {
		fmt.Printf("error decoding input file: %s\n", err.Error())
		return
	}

	// Iterate over the slice and display each user value.
	for _, u := range users {
		fmt.Printf("%v\n", u)
		// Marshal each user value and display the JSON. Check for errors.
		m, err := json.Marshal(&u)
		if err != nil {
			fmt.Printf("error opening marshalling user: %s\n", err.Error())
		}

		fmt.Println(string(m))
	}

}
