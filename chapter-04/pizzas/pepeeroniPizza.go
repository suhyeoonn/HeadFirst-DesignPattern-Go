package pizzas

import (
	"fmt"
)

type PepperoniPizza struct{}

// bake implements Pizza.
func (p PepperoniPizza) Bake() {
	fmt.Println("unimplemented")
}

// box implements Pizza.
func (p PepperoniPizza) Box() {
	fmt.Println("unimplemented")
}

// cut implements Pizza.
func (p PepperoniPizza) Cut() {
	fmt.Println("unimplemented")
}

// prepare implements Pizza.
func (p PepperoniPizza) Prepare() {
	fmt.Println("unimplemented")
}
