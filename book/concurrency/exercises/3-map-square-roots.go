package exercises

import (
	"fmt"
	"math"
	"sync"
)

// Write a function that builds a map[int]float64 where the keys are the numbers from 0 (inclusive)
// to 100,000 (exclusive) and the values are the square roots of those numbers (use the math.Sqrt function
// to calculate square roots). Use sync.OnceValue to generate a function that caches the map returned
// by this function and use the cached value to look up square roots for every 1,000th number from 0 to 100,000

type CustomMap map[int]float64

var sqrtMapCached func() CustomMap = sync.OnceValue(buildSqrtMap)

func buildSqrtMap() CustomMap {
	sqrtMap := make(CustomMap, 100000)

	for i := 0; i < 100000; i++ {
		sqrtMap[i] = math.Sqrt(float64(i))
	}

	return sqrtMap
}

func MapLookup() {
	for i := 0; i < 100000; i += 1000 {
		fmt.Println(sqrtMapCached()[i])
	}
}
