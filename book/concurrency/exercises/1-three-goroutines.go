package exercises

import (
	"context"
	"fmt"
	"sync"
)

// Create a function that launches three goroutines that communicate using a channel.
// The first two goroutines each write 10 numbers to the channel. The third goroutine reads all the numbers
// from the channel and prints them out. The function should exit when all values have been printed out.
// Make sure that none of the goroutines leak. You can create additional goroutines if needed

func ThreeGoroutines_Context_WaitGroup() { // the right Solution
	ch := make(chan string, 20)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				select {
				case <-ctx.Done(): // ensure goroutine won't leak
					return
				case ch <- fmt.Sprintf("channel #%d: %d", i, i*j):
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func ThreeGoroutines_Context_WaitGroup_Select() { // the right Solution
	ch := make(chan string, 20)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				select {
				case <-ctx.Done(): // ensure goroutine won't leak
					return
				case ch <- fmt.Sprintf("channel #%d: %d", i, i*j):
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

func ThreeGoroutines_Primitive() {
	ch := make(chan string, 20)

	for i := 1; i <= 2; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				ch <- fmt.Sprintf("channel #%d: %d", i, i*j)
			}
		}()
	}

	for j := 1; j <= 20; j++ { // can make a deadlock just by increasing the counter to 21
		fmt.Println(<-ch) // can leak goroutines by breaking the loop in the middle
	}
}

func ThreeGoroutines_Context() {
	ch := make(chan string, 20)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= 2; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				select {
				case <-ctx.Done():
					return
				case ch <- fmt.Sprintf("channel #%d: %d", i, i*j):
				}
			}
		}()
	}

	for j := 1; j <= 20; j++ { // can make a deadlock just by increasing the counter to 21
		fmt.Println(<-ch)
	}
}
