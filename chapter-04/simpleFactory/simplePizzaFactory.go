package simpleFactory

import (
	"chapter-04/pizzas"
)

type SimplePizzaFactory struct{}

func (s *SimplePizzaFactory) createPizza(pizzaType string) (pizza pizzas.Pizza) {
	if pizzaType == "cheese" {
		pizza = pizzas.CheesePizza{}
	} else if pizzaType == "greek" {
		pizza = pizzas.GreekPizza{}
	} else if pizzaType == "pepperoni" {
		pizza = pizzas.PepperoniPizza{}
	}
	return
}
