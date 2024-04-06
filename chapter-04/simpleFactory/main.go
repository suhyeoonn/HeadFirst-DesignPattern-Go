package main

import (
	"design-pattern/chapter-04/simpleFactory/chicagoPizza"
	"design-pattern/chapter-04/simpleFactory/nyPizza"
)

func main() {
	nyFactory := nyPizza.NewNYPizzaFactory()
	nyStore := nyPizza.NewPizzaStore(nyFactory)
	nyStore.OrderPizza("cheese")

	chicagoFactory := chicagoPizza.NewChicagoFactory()
	chicagoStore := chicagoPizza.NewPizzaStore(chicagoFactory)
	chicagoStore.OrderPizza("cheese")
}
