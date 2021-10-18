package main

import "fmt"

/*
goal : Find the optimal weight and value
sol:
- use bottom-up dp
- at each capacity, keep track of maximum weight

*/

type CakeType struct {
	weight int
	value  int
}

type CakeTypes []CakeType

func main() {

	caketype1 := &CakeType{7, 160}
	caketype2 := &CakeType{3, 90}
	caketype3 := &CakeType{2, 15}
	var caketypes []CakeType
	caketypes = append(caketypes, *caketype1, *caketype2, *caketype3)

	fmt.Println(maxDuffelBagValue(caketypes, 20))

}

func maxDuffelBagValue(cakeTypes CakeTypes, weightCapacity int) int {

	hashBestSoFar := make([]int, weightCapacity+1)
	//hashBestSoFar := make(map[int]int)
	for currentCapacity := 0; currentCapacity <= weightCapacity; currentCapacity++ {

		currentMax := 0

		for _, caketype := range cakeTypes {

			if caketype.weight <= currentCapacity {

				maxValueUsingCake := caketype.value + hashBestSoFar[currentCapacity-caketype.weight]

				currentMax = Max(currentMax, maxValueUsingCake)

			}
		}

		hashBestSoFar[currentCapacity] = currentMax
	}

	return hashBestSoFar[weightCapacity]

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
