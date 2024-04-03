package pizzas

type Pizza interface {
	prepare()
	bake()
	cut()
	box()
}
