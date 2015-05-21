// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// Package toy contains support for managing toy inventory.
package toy

// Declare a struct type named Toy with four fields. Name string,
// Weight int, onHand int and sold int.
type Toy struct {
	Name string
	Weight int
	onHand int
	sold int
}

// Declare a function named New that accepts values for the
// exported fields. Return a pointer of type Toy that is initialized
// with the parameters.
func NewToy(name string, weight int) *Toy {
	return &Toy{name, weight,0,0}
}

// Declare a method named OnHand with a pointer receiver that
// returns the current on hand count.
func (t *Toy) OnHand() int {
	return t.onHand
}

// Declare a method named UpdateOnHand with a pointer receiver that
// updates and returns the current on hand count.
func (t *Toy) UpdateOnHand(oh int) {
	t.onHand = oh
}

// Declare a method named Sold with a pointer receiver that
// returns the current sold count.
func (t *Toy) Sold() int {
	return t.sold
}

// Declare a method named UpdateSold with a pointer receiver that
// updates and returns the current sold count.
func (t *Toy)UpdateSold(s int) {
	t.sold = s
}