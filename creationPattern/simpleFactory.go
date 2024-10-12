package creationpattern

import "errors"

// if no factory pattern
// service logic -> basic struct mode
// if use factory pattern
// service logic -> factory pattern -> basic struct mode
// shortcoming: violate open-close principle

// abstract struct
type animal interface {
	eat()
	sleep()
}
type animalType int

const (
	_people = iota
	_duck
)

// implementation
type people struct {
}

func (p *people) eat()   {}
func (p *people) sleep() {}

type duck struct{}

func (d *duck) eat()   {}
func (d *duck) sleep() {}

type simpleFactory struct{}

func (sf *simpleFactory) NewAninal(t animalType) (animal, error) {
	switch t {
	case _people:
		return &people{}, nil
	case _duck:
		return &duck{}, nil
	default:
		return nil, errors.New("not supported animal type")
	}
}
