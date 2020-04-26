package Facade

import "fmt"

type CarModel struct {
}

func CreateModel() *CarModel {
	return &CarModel{}
}

func (car *CarModel) SetModel() {
	fmt.Println("Set-CarModel")
}

type CarEngine struct {
}

func CreateEngine() *CarEngine {
	return &CarEngine{}
}

func (car *CarEngine) SetModel() {
	fmt.Println("Set-CarEngine")
}

type CarBody struct {
}

func CreateBody() *CarBody {
	return &CarBody{}
}

func (car *CarBody) SetModel() {
	fmt.Println("Set-CarBody")
}

type CarFacade struct {
	model  CarModel
	engine CarEngine
	body   CarBody
}

func Init() *CarFacade {
	return &CarFacade{
		model:  CarModel{},
		engine: CarEngine{},
		body:   CarBody{},
	}
}

func (c *CarFacade) Create() {
	c.model.SetModel()
	c.engine.SetModel()
	c.body.SetModel()
}
