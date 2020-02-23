package main

import "fmt"

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}

	output := intersection(nums1, nums2)

	fmt.Println(output)

}

func intersection(nums1 []int, nums2 []int) []int {

	dict := make(map[int]bool)
	for _, num := range nums1 {
		dict[num] = true
	}

	var output []int
	for _, num := range nums2 {

		val, _ := dict[num]
		if val {
			output = append(output, num)
			dict[num] = false
		}
	}

	return output
}
