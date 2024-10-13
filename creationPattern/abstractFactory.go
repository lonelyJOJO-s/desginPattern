package creationpattern

// abstract factory pattern
// solve the 1:1 (factory:struct)

// 产品族， 产品等级结构

// 抽象层
type abstractApple interface {
	show()
}
type abstractBanana interface {
	show()
}
type abstractPear interface {
	show()
}

type abstractFactory interface {
	CreateApple() abstractApple
	CreateBanana() abstractBanana
	CreatePear() abstractPear
}

// 实现层
// 中国产品族
type chinaApple struct{}

func (ca *chinaApple) show() {}

type chinaBanana struct{}

func (cb *chinaBanana) show() {}

type chinaPear struct{}

func (cp *chinaPear) show() {}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() abstractApple {
	return &chinaApple{}
}
func (cf *ChinaFactory) CreateBanana() abstractBanana {
	return &chinaBanana{}
}
func (cf *ChinaFactory) CreatePear() abstractPear {
	return &chinaPear{}
}

// others 产品族
