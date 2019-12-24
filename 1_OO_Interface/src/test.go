package main

import (
	"fmt"
	"pet/struct"
	. "zoo/service"
)

func main() {
	var cat = _struct.Cat{Animal: Animal{
		Name:      "thomas",
		FootCount: 4,
		Food: "cat can",
		Toy: "mouse",
	}}
	var dog = _struct.Dog{Animal: Animal{
		Name:      "bruce",
		FootCount: 4,
		Food:      "dog can",
		Toy:       "ball",
	}}
	fmt.Printf("%T\n", cat)
	fmt.Println(Animaler(cat).Eat())
	fmt.Println(Animaler(cat).Play())
	fmt.Println()

	fmt.Printf("%T\n", dog)
	fmt.Println(Animaler(dog).Eat())
	fmt.Println(Animaler(dog).Play())
	fmt.Println()
}
