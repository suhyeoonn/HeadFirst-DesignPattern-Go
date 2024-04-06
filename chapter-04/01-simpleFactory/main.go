package main

import (
	"design-pattern/chapter-04/01-simpleFactory/chicagoPizza"
	"design-pattern/chapter-04/01-simpleFactory/nyPizza"
	"fmt"
)

func main() {
	fmt.Println("=====NY Pizza=====")
	nyFactory := nyPizza.NewNYPizzaFactory()
	nyStore := nyPizza.NewPizzaStore(nyFactory)
	nyStore.OrderPizza("cheese")

	fmt.Println("=====Chicago Pizza=====")
	chicagoFactory := chicagoPizza.NewChicagoFactory()
	chicagoStore := chicagoPizza.NewPizzaStore(chicagoFactory)
	chicagoStore.OrderPizza("cheese")
	// 시카고 지점에서는 종종 피자를 자르지 않는 실수가 발생!
	// OrderPizza 메소드에 cut이 주석처리 됨
}
