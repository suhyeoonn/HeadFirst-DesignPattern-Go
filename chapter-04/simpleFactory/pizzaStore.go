package simpleFactory

import "chapter-04/pizzas"

type PizzaStore struct {
	factory SimplePizzaFactory
}

func new(factory SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (p *PizzaStore) orderPizza(pizzaType string) (pizza pizzas.Pizza) {

	pizza = p.factory.createPizza(pizzaType)

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return
}
