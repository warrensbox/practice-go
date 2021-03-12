package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(allPossibleKFactors(2))
}

func allPossibleKFactors(k int) []float64 {
	values := []float64{}

	for a := 0; a <= k; a++ {
		powA := math.Pow(3, float64(a))
		for b := 0; b <= k; b++ {
			powB := math.Pow(5, float64(b))
			for c := 0; c <= k; c++ {
				powC := math.Pow(7, float64(c))
				//fmt.Printf("a:%v, b:%v, c:%v\n", a, b, c)
				value := powA * powB * powC
				fmt.Printf("value: %v\n", value)

				values = append(values, value)
			}
		}
	}
	sort.Float64s(values)
	return values
}
