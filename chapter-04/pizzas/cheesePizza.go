package pizzas

import "fmt"

type CheesePizza struct{}

func New() *CheesePizza {
	return &CheesePizza{}
}

// bake implements Pizza.
func (c CheesePizza) Bake() {
	fmt.Println("얇은 빵에 치즈 조금")
}

// box implements Pizza.
func (c CheesePizza) Box() {
	fmt.Println("unimplemented")
}

// cut implements Pizza.
func (c CheesePizza) Cut() {
	fmt.Println("unimplemented")
}

// prepare implements Pizza.
func (c CheesePizza) Prepare() {
	fmt.Println("unimplemented")
}
