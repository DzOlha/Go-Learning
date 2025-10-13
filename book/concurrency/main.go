package main

import (
	"book-concurrency/exercises"
	"fmt"
)

func main() {
	fmt.Println("---------------------------------Task #1: Solution 1------------------------------------")
	exercises.ThreeGoroutines_Context_WaitGroup() // solution #1 (use for-range)
	fmt.Println("---------------------------------Task #1: Solution 2------------------------------------")
	exercises.ThreeGoroutines_Context_WaitGroup_Select() // solution #2 (use for-select)

	fmt.Println("---------------------------------Task #2: Solution------------------------------------")
	exercises.TwoGoroutines()

	fmt.Println("---------------------------------Task #3: Solution------------------------------------")
	exercises.MapLookup()
}
