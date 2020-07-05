package service

import "pet"


func AnimalFactory(name string, animal Animal) Animaler {
	var animaler Animaler
	switch name {
		case "cat":
			animaler = &pet.Cat{Animal: &animal}
		case "dog":
			animaler = &pet.Dog{Animal: &animal}
		default:
			animaler = new(pet.Cat)
	}
	return animaler
}
