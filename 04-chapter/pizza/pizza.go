package pizza

import "fmt"

type Pizza struct {
	Name     string
	Dough    string
	Sauce    string
	Toppings []string
}

func (p *Pizza) Prepare() {
	fmt.Println("🍕 준비 중: ", p.Name)
	fmt.Println("🫓 도우를 돌리는 중...")
	fmt.Println("🥫 소스를 뿌리는 중...")
	fmt.Println("🧀 토핑을 올리는 중...")
	for _, topping := range p.Toppings {
		fmt.Println(" -", topping)
	}
}

func (p *Pizza) Bake() {
	fmt.Println("🔥 175도에서 25분 간 굽기")
}

func (p *Pizza) Cut() {
	fmt.Println("🔪 피자를 사선으로 자르기")
}

func (p *Pizza) Box() {
	fmt.Println("🎁 상자에 피자 담기")
}
