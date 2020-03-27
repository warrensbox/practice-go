package main

import "fmt"

import "math"

import "sort"

func main() {
	//points := [][]int{{1, 3}, {-2, 2}}
	points := [][]int{{0, 1}, {1, 0}}
	//points := [][]int{{10, 10}, {3, 3}, {5, -1}, {-2, 4}}

	//K := 1
	K := 2

	output := kClosest(points, K)

	fmt.Println(output)
}

func kClosest(points [][]int, K int) [][]int {

	if K == 0 {
		return [][]int{}
	}

	var result [][]int

	// create a Map with the distance of each point as key
	// and the index of the point in points as value
	dict := make(map[float64][]int)
	//dict := make(map[float64]int)
	for k, val := range points {

		dist := getDistance(val[0], val[1])
		dict[dist] = append(dict[dist], k)
		//dict[dist] = k
	}

	fmt.Println("MAP", dict)

	// get all distances and sort them
	keys := make([]float64, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	// get the K smallest distances and append their points to the result list
	found := 0
	for _, k := range keys {

		if K <= found {
			break
		}

		found += len(dict[k])

		for _, i := range dict[k] {
			result = append(result, points[i])
		}
	}

	return result

}

func getDistance(a, b int) float64 {

	dist := math.Sqrt(float64(a*a + b*b))

	return dist
}
