package _6_interfaces_generic_code

// we define generic code which will work with any type we pass
// type is defined at the moment we pass concrete value
func add[T int | float64 | string, K any](a, b T) K {
	return a + b
}

// add(1, 2) -> type will be int
