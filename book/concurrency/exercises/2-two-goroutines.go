package exercises

import (
	"context"
	"fmt"
)

// Create a function that launches two goroutines. Each goroutine writes 10 numbers to its own channel.
// Use a for-select loop to read from the both channels, printing out the number
// and the goroutine that wrote the value. Make sure that your function exits after all values are read
// and that none of your goroutines leak

func TwoGoroutines() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		defer close(ch1)
		for j := 1; j <= 10; j++ {
			select {
			case <-ctx.Done(): // ensure goroutine won't leak
				return
			case ch1 <- j:
			}
		}
	}()

	go func() {
		defer close(ch2)
		for j := 10; j <= 20; j++ {
			select {
			case <-ctx.Done(): // ensure goroutine won't leak
				return
			case ch2 <- j:
			}
		}
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case v1, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Printf("channel #1: %d \n", v1)
		case v2, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Printf("channel #2: %d \n", v2)
		}
	}
}
