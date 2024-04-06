package pizzas

import (
	"fmt"
)

type GreekPizza struct{}

// bake implements Pizza.
func (g GreekPizza) Bake() {
	fmt.Println("unimplemented")
}

// box implements Pizza.
func (g GreekPizza) Box() {
	fmt.Println("unimplemented")
}

// cut implements Pizza.
func (g GreekPizza) Cut() {
	fmt.Println("unimplemented")
}

// prepare implements Pizza.
func (g GreekPizza) Prepare() {
	fmt.Println("unimplemented")
}
