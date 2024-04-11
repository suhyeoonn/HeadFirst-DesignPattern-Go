package main

import (
	chicagostore "design-pattern/04-chapter/02-factoryMethod/chicagoStore"
	"design-pattern/04-chapter/02-factoryMethod/newYorkStore"
	"fmt"
)

func main() {
	fmt.Println("=====NY Pizza=====")
	nyPizzaStore := &newYorkStore.NYPizzaStore{}
	nyPizzaStore.PizzaStore.CreatePizza = nyPizzaStore.CreatePizza
	nyPizzaStore.OrderPizza("cheese")

	fmt.Println("=====Chicago Pizza=====")
	chicagoStore := chicagostore.New()
	chicagoStore.OrderPizza("cheese")
}
