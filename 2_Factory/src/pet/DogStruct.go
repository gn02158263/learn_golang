package pet

import (
	zoo "zoo/service"
)

type Dog struct {
	*zoo.Animal
}

func (dog *Dog) Eat() string {
	return dog.Name +" eat the "+ dog.Food
}

func (dog *Dog) Play() string {
	return dog.Name +" play the "+ dog.Toy
}

