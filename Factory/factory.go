package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type Donglaishun struct {
}

func (d *Donglaishun) GetFood() {
	fmt.Println("来东来顺吃饭")
}

type Quanjude struct {
}

func (d *Quanjude) GetFood() {
	fmt.Println("全聚德也成啊")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Quanjude{}
	}
	return nil
}
