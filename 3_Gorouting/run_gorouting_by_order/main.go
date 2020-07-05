package main

import (
	"fmt"
	"time"
)

type Message struct {
	Text string
	Wait chan bool
}

func mergeChannel(a, b <-chan Message) <-chan Message {
	c := make(chan Message)
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
	aChannel := make(chan Message)
	bChannel := make(chan Message)
	fmt.Println("start")
	go addJob("a", aChannel)
	go addJob("b", bChannel)
	channel := mergeChannel(aChannel, bChannel)
	for i := 0; i <= 10; i++ {
		msg1 := <-channel
		msg2 := <-channel
		msg1.Wait <-true
		msg2.Wait <-true
		fmt.Println("say:", msg1.Text)
		fmt.Println("say:", msg2.Text)
	}
	fmt.Println("end")
}

func addJob(msg string, channel chan Message) {
	waitForIt := make(chan bool)
	for i := 0; ; i++ {
		channel <- Message{fmt.Sprintf("%d %s", i, msg), waitForIt}
		time.Sleep(time.Second)
		<-waitForIt
	}
}
