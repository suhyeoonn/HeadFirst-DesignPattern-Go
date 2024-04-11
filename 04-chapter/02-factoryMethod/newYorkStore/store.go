package newYorkStore

import (
	pizzastore "design-pattern/04-chapter/02-factoryMethod/pizzaStore"
	"design-pattern/04-chapter/pizza"
)

type NYPizzaStore struct {
	pizzastore.PizzaStore
}

func (s *NYPizzaStore) CreatePizza(pizzaType string) (pizza pizzastore.IPizza) {
	if pizzaType == "cheese" {
		pizza = NewNYCheesePizza()
	} else if pizzaType == "greek" {
		pizza = &greekPizza{}
	} else if pizzaType == "pepperoni" {
		pizza = &pepperoniPizza{}
	}
	return
}

func New() *NYPizzaStore {
	return &NYPizzaStore{}
}

type NYCheesePizza struct {
	pizza.Pizza
}

func NewNYCheesePizza() *NYCheesePizza {
	return &NYCheesePizza{
		Pizza: pizza.Pizza{
			Name:     "뉴욕 스타일 소스와 치즈 피자",
			Dough:    "씬 크러스트 도우",
			Sauce:    "마리나라 소스",
			Toppings: []string{"잘게 썬 레지아노 치즈"},
		},
	}
}

type greekPizza struct {
	pizza.Pizza
}

type pepperoniPizza struct {
	pizza.Pizza
}
