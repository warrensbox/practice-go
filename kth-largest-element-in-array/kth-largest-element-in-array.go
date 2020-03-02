package main

import (
	"fmt"
)

func main() {

	input := []int{3, 2, 1, 5, 5, 6, 6, 4}
	k := 2

	output := findKthLargest(input, k)
	fmt.Println(output)
}

func findKthLargest(nums []int, k int) int {

	return doFindKthLargest(nums, k, 0, len(nums)-1)

}

func doFindKthLargest(nums []int, k int, start int, end int) int {

	targetPos := len(nums) - k
	fmt.Println("targetPos", targetPos)
	fmt.Println("NUMS", nums)
	for {
		fmt.Println("start", start)
		fmt.Println("end", end)
		fmt.Println("----")

		pivotIndex := partition(nums, start, end)
		fmt.Println("pivotIndex", pivotIndex)
		if pivotIndex == targetPos {
			return nums[pivotIndex]
		} else if pivotIndex < targetPos {
			start = pivotIndex + 1
		} else {
			end = pivotIndex - 1
		}

	}

}

// partition choses nums[start] as pivot and
// moves all values less than pivot to the left side.
// moves all values greater then pinot to the right side
// of the pivot value
func partition(nums []int, start int, end int) int {
	fmt.Println("-00000-")
	pivot := nums[start]
	fmt.Println("pivot", pivot)
	left := start + 1
	right := end
	fmt.Println("left", left)
	fmt.Println("right", right)

	for left <= right {
		for left <= right && nums[left] <= pivot {
			fmt.Println("nums[left]", nums[left])
			fmt.Println("left before", left)
			fmt.Println("NUMS", nums)
			left++
			fmt.Println("left", left)
			fmt.Println("-LLLLLLL-")
		}
		for left <= right && nums[right] > pivot {
			fmt.Println("nums[right]", nums[right])
			fmt.Println("right", right)
			fmt.Println("NUMS", nums)
			right--
			fmt.Println("right", right)
			fmt.Println("-RRRRR-")
		}
		fmt.Println("-FFFFF-")
		fmt.Println("left", left)
		fmt.Println("right", right)
		fmt.Println("NUMS", nums)
		if left <= right {
			fmt.Println("<<<<", nums)
			nums[left], nums[right] = nums[right], nums[left]
		}
		fmt.Println("NUMS", nums)
	}
	fmt.Println("SWITCING", nums)
	nums[right], nums[start] = nums[start], nums[right]
	fmt.Println("NUMS", nums)
	return right
}
