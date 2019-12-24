package _struct

import "zoo/service"

type Cat struct {
	service.Animal
}

func (cat Cat) Eat() string {
	return cat.Name +" eat the "+ cat.Food
}

func (cat Cat) Play() string {
	return cat.Name +" play the "+ cat.Toy
}

