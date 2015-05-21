// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/9_MSdcdlNQ

// Declare a struct type named Point with two fields, X and Y of type int.
// Implement a factory function for this type and a method that accept a value
// of this type and calculates the distance between the two points. What is
// the nature of this type?
package main

// Add imports.
import (
	"fmt"
	"math"
)

// Declare struct type named Point that represents a point in space.
type Point struct {
	x, y float64
}

// Declare a function named New that returns a Point based on X and Y
// positions on a graph.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Declare a method named Distance that finds the length of the hypotenuse
// between two points. Pass one point in and return the answer.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
// Use the math.Pow and math.Sqrt functions.
func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p.x, 2.0) + math.Pow(p2.y-p.y, 2.0))
}

// main is the entry point for the application.
func main() {
	// Declare the first point.
	p1 := NewPoint(0, 0)

	// Declare the second point.
	p2 := NewPoint(1, 1)

	// Calculate the distance and display the result.
	fmt.Printf("distance: %v\n", p1.Distance(p2))

}
