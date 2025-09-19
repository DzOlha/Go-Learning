package _7_arrays_slices_maps

import "fmt"

func setupPrices() {
	// [...] allows us to set the size using the number of items mentioned between {}
	prices := [...]float64{2.3, 4.7, 56.89, 89.76}
	fmt.Println(prices)

	// [4] size set explicitly
	prices2 := [4]float64{2.3, 4.7, 56.89, 89.76}
	fmt.Println(prices2)

	var prices3 = [2]string{"Hello", "World"}
	fmt.Println(prices3)

	var arr [2]string
	arr = [2]string{"Hello", "World"}
	fmt.Println(arr)
}
