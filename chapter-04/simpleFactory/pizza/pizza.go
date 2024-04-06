package pizza

type IPizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
