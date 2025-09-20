package _0_concurrency

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello!", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChannel <- true // write data to the channel

	// close the channel when we use single channel approach and if we know that the operation will take the longest
	close(doneChannel) // explicitly tell go when it can stop waiting for a new values from a channel (because it is closed)
}

func App() {
	multipleChannelApproach()
	//oneChannelApproach()
}

func multipleChannelApproach() {
	// we can use one and the same channel with multiple goroutines
	//done := make(chan bool) // channel is a communication/transmission device
	// it is intended to be used to receive multiple values from different goroutines

	dones := make([]chan bool, 4)
	dones[0] = make(chan bool)

	// running function in goroutine - run it in a non-blocking way (so the next function can be immediately invoked)
	go greet("Nice to meet you!", dones[0]) // dispatch goroutine

	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])

	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])

	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done // read data from the channel
	}
}

func oneChannelApproach() {
	// we can use one and the same channel with multiple goroutines
	//done := make(chan bool) // channel is a communication/transmission device
	// it is intended to be used to receive multiple values from different goroutines

	done := make(chan bool)

	// running function in goroutine - run it in a non-blocking way (so the next function can be immediately invoked)
	go greet("Nice to meet you!", done) // dispatch goroutine
	go greet("How are you?", done)      // goroutines do not return any values, so we need to use channels to transmit data
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {

	}
}
