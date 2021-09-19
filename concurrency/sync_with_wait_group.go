package main

import (
	"fmt"
	"sync"
)

func publish(publisherId int, messages chan<- string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Publisher", publisherId, "publishes", i)
		messages <- fmt.Sprintf("Message from publisher %v: %v", publisherId, i)
		//time.Sleep(10 * time.Millisecond)
	}
}

func consume(consumerId int, messages <-chan string) {

	for {
		m := <-messages
		if m == "" {
			fmt.Println("* *** Consumer", consumerId, "finished")

			return
		}
		fmt.Println("* Consumer", consumerId, m)
	}
}

func main() {

	var waitForPublishers, waitForConsumers sync.WaitGroup
	messages := make(chan string)

	waitForPublishers.Add(1)
	go func() {
		publish(1, messages, 100)
		waitForPublishers.Done()
	}()

	waitForPublishers.Add(1)
	go func() {
		publish(2, messages, 100)
		waitForPublishers.Done()
	}()

	waitForConsumers.Add(1)
	go func() {
		consume(1, messages)
		waitForConsumers.Done()
	}()

	waitForConsumers.Add(1)
	go func() {
		consume(2, messages)
		waitForConsumers.Done()
	}()

	waitForConsumers.Add(1)
	go func() {
		consume(3, messages)
		waitForConsumers.Done()
	}()

	waitForPublishers.Wait()
	close(messages)
	waitForConsumers.Wait()

	fmt.Println("Fin!")
}
