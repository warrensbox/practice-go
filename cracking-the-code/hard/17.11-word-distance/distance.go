package main

import (
	"fmt"
	"math"
)

func main() {
	// converter()words := []string{""}
	words := converter("Many of the attendees arrived festooned in purple the official color of the Rouhanicampaign purple signs and polo shirts purple ribbons wrapped around wrists and fingers")
	findClosest(words, "Many", "purple")
}

func findClosest(words []string, word1, word2 string) {
	hash := make(map[string][]int)
	getWordLocations(words, hash)
	//fmt.Println(hash)
	min := findMinDistancePair(word1, word2, hash)
	fmt.Println("min", min)
}

func getWordLocations(words []string, hash map[string][]int) {

	for index, val := range words {
		hash[val] = append(hash[val], index)
	}
}

func findMinDistancePair(word1, word2 string, hash map[string][]int) int {

	index1 := 0
	index2 := 0
	min := math.MaxInt16
	//hash get array of positions for wordX
	arr1 := hash[word1]
	arr2 := hash[word2]
	for index1 < len(arr1) && index2 < len(arr2) {
		distance := Abs(arr1[index1] - arr2[index2])
		if distance < min {
			min = distance
		}
		if arr1[index1] < arr2[index2] {
			index1++
		} else {
			index2++
		}
	}

	return min
}

/* HELPER FUNCTION - CAN BE IGNORED */
/* compose sentence */
func split(tosplit string, sep rune) []string {
	var fields []string

	last := 0
	for i, c := range tosplit {
		if c == sep {
			// Found the separator, append a slice
			fields = append(fields, string(tosplit[last:i]))
			last = i + 1
		}
	}

	// Don't forget the last field
	fields = append(fields, string(tosplit[last:]))

	return fields
}

func converter(str string) []string {

	arr := []string{}
	for _, field := range split(str, ' ') {
		arr = append(arr, field)
	}

	return arr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
