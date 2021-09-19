package main

import (
	"fmt"
)

func publish(publisherId int, messages chan<- string, count int, done chan<- bool) {
	for i := 0; i < count; i++ {
		fmt.Println("Publisher", publisherId, "publishes", i)
		messages <- fmt.Sprintf("Message from publisher %v: %v", publisherId, i)
		//time.Sleep(10 * time.Millisecond)
	}

	done <- true
}

func consume(consumerId int, messages <-chan string, done chan<- bool) {

	for {
		m := <-messages
		if m == "" {
			fmt.Println("* *** Consumer", consumerId, "finished")

			done <- true
			return
		}
		fmt.Println("* Consumer", consumerId, m)
	}
}

func main() {
	messages := make(chan string)

	publisher1Done := make(chan bool, 1)
	go publish(1, messages, 100, publisher1Done)

	publisher2Done := make(chan bool, 1)
	go publish(2, messages, 100, publisher2Done)

	consumer1Done := make(chan bool, 1)
	go consume(1, messages, consumer1Done)

	consumer2Done := make(chan bool, 1)
	go consume(2, messages, consumer2Done)

	consumer3Done := make(chan bool, 1)
	go consume(3, messages, consumer3Done)

	// Wait for the publishers to finish.
	// When publishers finsished, close the channel.
	<-publisher1Done
	<-publisher2Done
	close(messages)

	// Wait for consumers to finish.
	<-consumer1Done
	<-consumer2Done
	<-consumer3Done

	fmt.Println("Fin!")
}
