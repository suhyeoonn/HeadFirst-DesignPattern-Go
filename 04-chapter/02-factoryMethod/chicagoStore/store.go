package chicagostore

import (
	"design-pattern/04-chapter/02-factoryMethod/pizzaStore"
	"design-pattern/04-chapter/pizza"
	"fmt"
)

type ChicagoPizzaStore struct {
	pizzaStore.PizzaStore
}

type chicagoCheesePizza struct {
	pizza.Pizza
}

type greekPizza struct {
	pizza.Pizza
}

type pepperoniPizza struct {
	pizza.Pizza
}

func (s *ChicagoPizzaStore) CreatePizza(pizzaType string) (pizza pizzaStore.IPizza) {
	if pizzaType == "cheese" {
		pizza = NewChicagoCheesePizza()
	} else if pizzaType == "greek" {
		pizza = &greekPizza{}
	} else if pizzaType == "pepperoni" {
		pizza = &pepperoniPizza{}
	}
	return
}

func New() *ChicagoPizzaStore {
	store := &ChicagoPizzaStore{}
	store.PizzaStore.CreatePizza = store.CreatePizza
	return store
}

func NewChicagoCheesePizza() *chicagoCheesePizza {
	return &chicagoCheesePizza{
		Pizza: pizza.Pizza{
			Name:     "시카고 스타일 딥 디쉬 치즈 피자",
			Dough:    "아주 두꺼운 크러스트 도우",
			Sauce:    "플럼토마토 소스",
			Toppings: []string{"잘게 조각낸 모짜렐라 치즈"},
		},
	}
}

func (c *chicagoCheesePizza) Cut() {
	fmt.Println("✂️ 네모난 모양으로 피자 자르기")
}
