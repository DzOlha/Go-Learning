package _8_functions

import "fmt"

func ExampleVariadicFunction() {
	sum := sumVariadic(54, 1, 6, 15)
	fmt.Println(sum)

	numbers := []int{14, 15, 18, 67}

	// convenient way to pass slice values without the need to mention each of them explicitly by index
	// we can unwrap slice to pass all of its items separately using ... operator after the slice name
	anotherSum := sumVariadic(58, numbers...)
	fmt.Println(anotherSum)
}

// such flexible number of parameters can be useful in case we do not want to create a slice
// just to pass it into a function, but still need to pass a lot of values of the same type into it
func sumVariadic(startingNum int, followingParams ...int) int {
	sum := startingNum
	size := len(followingParams) // behind the scene sequence of such params are collected into a slice

	for i := 0; i < size; i++ {
		sum += followingParams[i]
	}

	return sum
}
