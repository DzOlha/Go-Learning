package _8_functions

import "fmt"

type transformFn func(int) int // create alias for function type to keep declaration short

func ExampleFunctionAsParameter() {
	numbers := []int{1, 2, 3, 4, 5}

	doubled := transformNumbers(&numbers, getTransformerFunction(&numbers)) // get function as a return value of another function
	tripled := transformNumbers(&numbers, triple)                           // pass function via its name as a parameter

	fmt.Println(doubled)
	fmt.Println(tripled)

	// pass anonymous function
	anonymouslyTransformedNumbers := transformNumbers(&numbers, func(num int) int {
		return num * 18
	})
	fmt.Println(anonymouslyTransformedNumbers)

	// define functions using transformer with closure under the hood
	doubleTransformer := createTransformer(2)
	tripleTransformer := createTransformer(3)

	doubled2 := transformNumbers(&numbers, doubleTransformer)
	tripled2 := transformNumbers(&numbers, tripleTransformer)
	fmt.Println(doubled2)
	fmt.Println(tripled2)

}

func transformNumbers(nums *[]int, transform transformFn) []int {
	var result []int

	for _, num := range *nums {
		result = append(result, transform(num))
	}

	return result
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

// function can return another function
func getTransformerFunction(nums *[]int) transformFn {
	if (*nums)[0] == 0 {
		return double
	} else {
		return triple
	}
}

// create a closure using anonymous function which locked in outer scope
func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}
