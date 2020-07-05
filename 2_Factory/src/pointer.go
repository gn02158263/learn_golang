package main

import "fmt"

type Service struct {
	data string
}

func main() {

	oldService := &Service{

		data: "test",
	}

	newService := *oldService
	newService2 := oldService

	newService.data = "111"

	fmt.Println("origin &oldService:", &oldService)
	fmt.Println("origin oldService:", oldService)
	fmt.Println("")

	fmt.Println("when newService := *oldService, &newService:", &newService)
	fmt.Println("when newService := *oldService, newService:", newService)
	fmt.Println("")

	fmt.Println("&oldService:", &oldService)
	fmt.Println("oldService:", oldService)
	fmt.Println("newService := *oldService is not effect oldService obj")

	fmt.Println("-----------------------------------------------------------")

	fmt.Println("origin &oldService:", &oldService)
	fmt.Println("origin oldService:", oldService)
	fmt.Println("")

	newService2.data = "222"

	fmt.Println("when newService2 := oldService, &newService2:", &newService2)
	fmt.Println("when newService2 := oldService, *newService2:", *newService2)
	fmt.Println("when newService2 := oldService, newService2:", newService2)
	fmt.Println("")

	fmt.Println("&oldService:", &oldService)
	fmt.Println("oldService:", oldService)
	fmt.Println("newService2 := oldService is will effect oldService obj")
}
