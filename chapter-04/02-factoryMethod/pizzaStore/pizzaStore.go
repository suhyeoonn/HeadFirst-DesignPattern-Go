package pizzaStore

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaStore struct {
	CreatePizza func(pizzaType string) IPizza
}

// 각 서브클래스는 createPizza 메소드를 오버라이드하지만,
// orderPizza 메소드는 여기서 정의한 내용 그대로 사용한다.
func (p *PizzaStore) OrderPizza(pizzaType string) IPizza {

	pizza := p.CreatePizza(pizzaType)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
