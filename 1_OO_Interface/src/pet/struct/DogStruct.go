package _struct

import "zoo/service"

type Dog struct {
	service.Animal
}

func (dog Dog) Eat() string {
	return dog.Name +" eat the "+ dog.Food
}

func (dog Dog) Play() string {
	return dog.Name +" play the "+ dog.Toy
}

