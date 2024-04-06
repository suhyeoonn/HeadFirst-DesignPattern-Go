package chicagoPizza

import (
	"design-pattern/chapter-04/01-simpleFactory/pizza"
	"fmt"
)

type ChicagoCheesePizza struct{}

func NewChicagoCheesePizza() *ChicagoCheesePizza {
	return &ChicagoCheesePizza{}
}

// bake implements Pizza.
func (c ChicagoCheesePizza) Bake() {
	fmt.Println("두꺼운 빵 위에 소스를 풍부하게")
}

// box implements Pizza.
func (c ChicagoCheesePizza) Box() {
	fmt.Println("Box")
}

// cut implements Pizza.
func (c ChicagoCheesePizza) Cut() {
	fmt.Println("Cut")
}

// prepare implements Pizza.
func (c ChicagoCheesePizza) Prepare() {
	fmt.Println("Prepare")
}

type ChicagoFactory struct{}

func (s *ChicagoFactory) createPizza(pizzaType string) (pizza pizza.IPizza) {
	if pizzaType == "cheese" {
		pizza = ChicagoCheesePizza{}
	}
	// else if pizzaType == "greek" {
	// 	pizza = pizzas.GreekPizza{}
	// } else if pizzaType == "pepperoni" {
	// 	pizza = pizzas.PepperoniPizza{}
	// }
	return
}

func NewChicagoFactory() *ChicagoFactory {
	return &ChicagoFactory{}
}

type pizzaStore struct {
	factory *ChicagoFactory
}

func NewPizzaStore(factory *ChicagoFactory) *pizzaStore {
	return &pizzaStore{factory: factory}
}

func (p *pizzaStore) OrderPizza(pizzaType string) (pizza pizza.IPizza) {

	pizza = p.factory.createPizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	// pizza.Cut()
	pizza.Box()

	return pizza
}
