package maps

import "fmt"

type floatMap map[string]float64 // type alias

func Maps() {
	courseRatings := map[string]float64{
		"Math":        64.7,
		"Programming": 67.9,
	}
	courseRatings["Arts"] = 45.7
	fmt.Println(courseRatings)

	courseRatings2 := make(floatMap, 3)
	courseRatings2["New"] = 45.7

	fmt.Println(courseRatings2)

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
