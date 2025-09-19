package _7_arrays_slices_maps

import (
	"fmt"
	"go-learning/main/07-arrays-slices-maps/product"
)

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func Solution() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	var hobbies = [3]string{"Reading", "Writing", "Coding"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	sublist := hobbies[1:]
	fmt.Println(sublist)

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1, slice2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	reslice := slice1[1:3]
	fmt.Println(reslice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"learn Go", "Practice Go"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Practice Go!!!"
	goals = append(goals, "Become the true Golang Engineer")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []*product.Product{
		product.New("Test title", 1, 34.6),
		product.New("Test title 2", 2, 38.6),
	}
	products = append(products, product.New("Test title 3", 3, 340.6))
	for _, p := range products {
		fmt.Printf("ID: %d, Title: %s, Price: %.2f\n", p.ID, p.Title, p.Price)
	}
}
