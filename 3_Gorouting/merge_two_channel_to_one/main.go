package main

import (
	"fmt"
	"time"
)


func mergeChannel(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-a
		}
	}()
	go func() {
		for {
			c <- <-b
		}
	}()

	return c
}

func main() {
	aChannel := make(chan string)
	bChannel := make(chan string)
	fmt.Println("start")
	go addJob("a", aChannel)
	go addJob("b", bChannel)
	channel := mergeChannel(aChannel, bChannel)
	for i := 0; i <= 10; i++ {
		fmt.Println("say:", <-channel)
	}
	fmt.Println("end")
}

func addJob(msg string, channel chan string) {
	for i := 0; ; i++ {
		channel <- fmt.Sprintf("%d %s", i, msg)
		time.Sleep(time.Second)
	}
}
