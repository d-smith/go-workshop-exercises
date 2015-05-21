// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initalize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import "fmt"



// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type ballplayer struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a batter.
func (b ballplayer) average() float64 {
	return float64(b.hits) / float64(b.atBats)
}

// main is the entry point for the application.
func main() {
	var players []ballplayer
	players = append(players, ballplayer{"a", 100, 30})
	players = append(players, ballplayer{"b", 1200, 10})
	players = append(players, ballplayer{"c", 200, 190})
	for _, p := range players {
		fmt.Printf("player %s hit %f\n", p.name, p.average())
	}
}
