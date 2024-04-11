package nyPizza

import (
	"design-pattern/04-chapter/01-simpleFactory/pizza"
	"fmt"
)

type NYPizzaFactory struct{}

func (s *NYPizzaFactory) createPizza(pizzaType string) (pizza pizza.IPizza) {
	if pizzaType == "cheese" {
		pizza = NYCheesePizza{}
	}
	// else if pizzaType == "greek" {
	// 	pizza = pizzas.GreekPizza{}
	// } else if pizzaType == "pepperoni" {
	// 	pizza = pizzas.PepperoniPizza{}
	// }
	return
}

func NewNYPizzaFactory() *NYPizzaFactory {
	return &NYPizzaFactory{}
}

type pizzaStore struct {
	factory *NYPizzaFactory
}

func NewPizzaStore(factory *NYPizzaFactory) *pizzaStore {
	return &pizzaStore{factory: factory}
}

func (p *pizzaStore) OrderPizza(pizzaType string) (pizza pizza.IPizza) {

	pizza = p.factory.createPizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}

type NYCheesePizza struct{}

func NewNYCheesePizza() *NYCheesePizza {
	return &NYCheesePizza{}
}

// bake implements Pizza.
func (c NYCheesePizza) Bake() {
	fmt.Println("얇은 빵에 치즈 조금")
}

// box implements Pizza.
func (c NYCheesePizza) Box() {
	fmt.Println("box")
}

// cut implements Pizza.
func (c NYCheesePizza) Cut() {
	fmt.Println("cut")
}

// prepare implements Pizza.
func (c NYCheesePizza) Prepare() {
	fmt.Println("Prepare")
}
