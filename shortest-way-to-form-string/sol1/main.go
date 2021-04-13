package main

import "fmt"

func main() {
	source := "abc"
	target := "abcbc"
	//source := "acdddpdab"
	//target := "apdd"

	output := shortestWay(source, target)

	fmt.Println("output", output)
}

func shortestWay(source string, target string) int {
	mapping := make(map[rune][]int)

	for i, val := range source {
		mapping[val] = append(mapping[val], i)
	}

	fmt.Println(mapping)

	idx := 0
	res := 1
	for _, val := range target {
		fmt.Println(string(val))
		if nums, ok := mapping[val]; !ok {
			return -1
		} else {
			if nums[len(nums)-1] < idx {
				fmt.Println("here")
				res++
				idx = 0
			}
			j := search(nums, idx)
			idx = nums[j] + 1
			fmt.Println("idx", idx)
		}
	}

	return res
}

func search(arr []int, val int) int {
	fmt.Println("val", val)
	lo := 0
	hi := len(arr) - 1
	res := -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == val {
			return mid
		} else if arr[mid] > val {
			res = mid
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return res
}
