# 팩토리 패턴
인터페이스에 맞춰서 코딩하면 시스템에서 일어날 수 있는 여러 변화에 대응할 수 있다. 인터페이스를 바탕으로 만들어진 코드는 어떤 클래스든 특정 인터페이스만 구현하면 사용할 수 있기 때문이다. => 다형성

반대로 구상 클래스가 추가될 때마다 코드를 고쳐야 하므로 수많은 문제가 생길 수 있다. => 변경에 닫혀 있는 코드

새로운 구상 형식을 써서 확장해야 할 때는 어떻게 해서든 다시 열 수 있게 만들어야 한다.

💡 첫 번째 디자인 원칙을 기억하자: 바뀌는 부분을 찾아내서 바뀌지 않는 부분과 분리해야 한다

## 바뀌는 부분 찾아내기
```go
func orderPizza(pizzaType string) (pizza Pizza){
	// 이 코드는 변경에 닫혀 있지 않다. 
	// 피자 종류가 바뀔 때마다 코드를 계속 고쳐야 한다.
	if pizzaType == "cheese" {
		pizza = new CheesePizza{}
	} else if pizzaType == "greek" {
		pizza = new GreekPizza{}
	} else if pizza == "pepperoni" {
		pizza = new PepperoniPizza{}
	}

	// 이 부분은 바뀌지 않는다. 피자를 판매할 때 당연히 해야 하는 일이다.
	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return
}
```

## 객체 생성 부분 캡슐화하기
객체 생성을 처리하는 클래스를 `팩토리(Factory)`라고 부른다. SimplePizza Factory를 만들면 이제 더 이상 orderPizza() 메소드에서 어떤 피자를 만들지 고민하지 않아도 된다. orderPizza() 메소드는 Pizza 인터페이스를 구현하는 피자를 받아서 그 인터페이스에서 정의했던 메소드를 호출하기만 하면 된다.

## 객체 생성 팩토리 만들기 (Simple Factory)
피자 객체 생성 부분을 전담할 팩토리를 정의하자.
```go
// 이 클래스가 하는 일은 단 하나다. 클라이언트가 받을 피자만 만든다.
type SimplePizzaFactory struct{}

func (s *SimplePizzaFactory) createPizza(pizzaType string) (pizza Pizza) {
	if pizzaType == "cheese" {
		pizza = CheesePizza{}
	} else if pizzaType == "greek" {
		pizza = GreekPizza{}
	} else if pizza == "pepperoni" {
		pizza = PepperoniPizza{}
	}
	return
}
```

클라이언트 코드 수정
```go
type PizzaStore struct {
	Pizza
	factory SimplePizzaFactory
}

func new(factory SimplePizzaFactory) *PizzaStore {
	return &PizzaStore{factory: factory}
}

func (p *PizzaStore) orderPizza(pizzaType string) (pizza Pizza) {

	pizza = p.factory.createPizza(pizzaType)

	pizza.prepare()
	pizza.bake()
	pizza.cut()
	pizza.box()
	return
}
```

### 🤔 이렇게 캡슐화하면 무슨 장점이 있는지?
- `SimplePizzaFactory`를 사용하는 클라이언트가 매우 많을 수 있다. 그런 상황에서 피자 객체 생성 작업을 팩토리 클래스로 캡슐화해 놓으면 구현을 변경할 때 여기저기 고칠 힐요 없이 팩토리 클래스 하나만 고치면 된다.

### 🤔 팩토리 정적 메소드와 차이점?
- 간단한 팩토리를 정적 메소드로 정의하는 기법도 많이 쓰이며 정적 패토리라고도 부른다. 정적 메소드를 쓰면 객체 생성 메소드를 실행하려고 *객체의 인스턴스를 만들지 않아도 되기 때문*이다. 하지만 서브클래스를 만들어서 객체 생성 메소드의 행동을 변경할 수 없다는 단점이 있다는 것도 기억하자.

## 팩토리 메소드 패턴
팩토리 메소드 패턴에서는 객체를 생성할 때 필요한 인터페이스를 만든다. 어떤 클래스의 인스턴스를 만들지는 서브클래스에서 결정한다. 팩토리 메소드 패턴을 사용하면 클래스 인스턴스 만드는 일을 서브클래스에게 맡기게 된다. (사용하는 서브클래스에 따라 생산되는 객체 인스턴스 결정)
### 생산자(Creator) 클래스
추상 생산자 클래스. 나중에 서브클래스에서 제품을 생산하려고 구현하는 팩토리 메소드를 정의한다. 생산자 자체는 어떤 구상 제품 클래스가 만들어질지 미리 알 수 없다.
```go
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

type NYPizzaStore struct {
	pizzastore.PizzaStore
}

// CreatePizza가 팩토리 메소드이다.
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
```
### 제품(Product) 클래스
팩토리는 제품을 생산한다.
```go
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
```
### 🤔 간단한 팩토리와 팩토리 메소드의 차이?
간단한 팩토리는 일회용 처방에 불과한 반면, 팩토리 메소드 패턴을 사용하면 여러 번 재사용이 가능한 프레임워크를 만들 수 있다.

팩토리 메소드 패턴의 `orderPizza()` 메소드는 피자를 만드는 일반적인 프레임워크를 제공한다. 모든 피자는 `orderPizza()`를 통해서만 만들어지기 피자 자르기를 깜빡하는 실수 등을 방지할 수 있다. 또한, 구상 클래스를 만들 때 `createPizza()` 추상 메소드가 정의되어 있는 추상 클래스를 확장해서 만들었다. `createPizza()` 메소드에서 어떤 일을 할지는 각 지점에서 결정한다. 이 프레임워크를 간단한 팩토리와 비교해보면 간단한 팩토리는 객체 생성을 캡슐화하는 방법을 사용하긴 하지만 팩토리 메소드만큼 유연하지는 않다. 팩토리가 PizzaStore 안에 포함되는 별개의 객체이며 생성하는 제품을 마음대로 변경할 수 없다.

간단한 팩토리는 PizzaStore에 속해 있는 팩토리 하나만 생각한다.
```go
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
```
시카고 지점이 필요하다면 pizzaStore를 또 다시 만들어야 했다. 이 과정에서 코드 중복이 발생하며, Order() 메소드에서 필요한 코드를 누락하는 실수도 발생할 수 있다.
```go
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
	// pizza.Cut() 누락!!
	pizza.Box()

	return pizza
}
```