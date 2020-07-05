package main

import (
	"fmt"
	zoo "zoo/service"
)

func getZooAnimal() *zoo.Animal {
	return new(zoo.Animal)
}

func main() {
	var cat = zoo.Animal{
		Name:      "thomas",
		FootCount: 4,
		Food: "cat can",
		Toy: "mouse",
	}
	var dog = zoo.Animal{
		Name:      "bruce",
		FootCount: 4,
		Food:      "dog can",
		Toy:       "ball",
	}
	var animaler = zoo.AnimalFactory("cat", &cat)

	fmt.Printf("%T\n", animaler)
	fmt.Println(animaler.Eat())
	fmt.Println(animaler.Play())
	fmt.Println()

	animaler = zoo.AnimalFactory("dog", &dog)
	fmt.Printf("%T\n", animaler)
	fmt.Println(animaler.Eat())
	fmt.Println(animaler.Play())
	fmt.Println()
}
