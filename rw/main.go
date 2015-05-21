// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/MxcJ581bt6

// Download any document from the web and display the content in
// the terminal and write it to a file at the same time.
package main

// Add imports.
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// main is the entry point for the application.
func main() {
	// Download the RSS feed for "http://www.goinggo.net/feeds/posts/default".
	// Check for errors.
	resp, err := http.Get("http://www.goinggo.net/feeds/posts/default")
	if err != nil {
		fmt.Printf("Error loading file: %s\n", err.Error())
		return
	}

	// Arrange for the response Body to be Closed using defer.
	defer resp.Body.Close()

	// Declare a slice of io.Writer interface values.
	var writers []io.Writer

	// Append stdout to the slice of writers.
	writers = append(writers, os.Stdout)

	// Open a file named "goinggo.rss" and check for errors.
	gfile, err := os.Create("goinggo.rss")
	if err != nil {
		fmt.Printf("Error loading file: %s\n", err.Error())
		return
	}

	// Close the file when the function returns.
	defer gfile.Close()

	// Append the file to the slice of writers.
	writers = append(writers, gfile)

	// Create a MultiWriter interface value from the writers
	// inside the slice of io.Writer values.
	dest := io.MultiWriter(writers...)

	// Write the response to both the stdout and file using the
	// MultiWriter. Check for errors.
	_, err = io.Copy(dest, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to multiwriter: %s\n", err.Error())
	}

}
