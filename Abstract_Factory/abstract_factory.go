package Abstract_Factory

import "fmt"

type Lunch interface {
	Cook()
}

type rise struct {
}

func (r *rise) Cook() {
	fmt.Println("焖米饭呢")
}

type Tomato struct {
}

func (t *Tomato) Cook() {
	fmt.Println("炒西红柿呢")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}

type SimpleLunchFactory struct {
}

func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}
