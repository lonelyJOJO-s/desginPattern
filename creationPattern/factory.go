package creationpattern

// factory pattern define a factory iface, which left each specific factory to implement
// slove the "violate open-close principle" shortcoming from simple factory pattern

type production interface {
	effect()
}

type car struct{}

func (c *car) effect() {}

type phone struct{}

func (p *phone) effect() {}

type factoryAbstract interface {
	createProduction() production
}

type carFactory struct{}

func (cf *carFactory) createProduction() production {
	fc := new(car)
	return fc
}

type phoneFactory struct{}

func (pf *phoneFactory) createProduction() production {
	fc := new(phone)
	return fc
}

func test() {
	carFac := new(carFactory)
	car := carFac.createProduction()
	car.effect()
	phoneFac := new(phoneFactory)
	phone := phoneFac.createProduction()
	phone.effect()
}
