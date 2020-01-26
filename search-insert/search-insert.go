package main

import (
	"fmt"
)

func main() {

	input := []int{1, 3, 5, 6}
	target := 7
	output := searchInsert(input, target)

	fmt.Println(output)
}

func searchInsert1(nums []int, target int) int {

	index, i := 0, 0
	for i < len(nums) {
		fmt.Println("nums", nums[i])

		if target-nums[i] == 0 {
			index = i
			break
		} else if target-nums[i] > 0 {
			if i == len(nums)-1 {
				i++
				index = i
				break
			} else {
				i++
			}
		} else if target-nums[i] < 0 {
			index = i
			break
		}
	}

	return index
}

// public int searchInsert(int[] A, int target) {
// 	int low = 0, high = A.length-1;
// 	while(low<=high){
// 		int mid = (low+high)/2;
// 		if(A[mid] == target) return mid;
// 		else if(A[mid] > target) high = mid-1;
// 		else low = mid+1;
// 	}
// 	return low;
// }

func searchInsert(nums []int, target int) int {

	low := 0
	high := len(nums) - 1

	for low <= high {

		fmt.Println("high", high)
		fmt.Println("low", low)
		fmt.Println("_______")
		mid := (low + high) / 2
		fmt.Println("mid", mid)
		fmt.Println("nums[mid]", nums[mid])

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
		fmt.Println("++++++++++")
	}

	fmt.Println("low", low)
	return low
}
