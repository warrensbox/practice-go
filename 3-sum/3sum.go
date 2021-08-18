package main

import (
	"fmt"
	"sort"
)

//popular question

func main() {

	input := []int{-5, -4, -4, -3, -3, -2, -1, -1, 0, 1, 2, 2, 3, 4, 4, 5, 6}

	//output := threeSum(input)
	output := sumOfTwo(input, 1, 0, len(input)-1)
	fmt.Println(output)
}

func threeSum(nums []int) [][]int {

	var res [][]int
	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	//fmt.Println("sorted", nums)

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//fmt.Println("sorted", nums)
		target := -nums[i]
		left := i + 1
		right := len(nums) - 1

		//fmt.Println("target", target)
		//fmt.Println("left", left)
		//fmt.Println("right", right)

		for left < right {

			//fmt.Println("nums[left]", nums[left])
			//fmt.Println("nums[right]", nums[right])

			if nums[left]+nums[right] == target {
				//fmt.Println("bingo")
				var inner []int

				inner = append(inner, -target)
				inner = append(inner, nums[left])
				inner = append(inner, nums[right])

				res = append(res, inner)

				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
					//fmt.Println("left", left)
					//fmt.Println("right", right)
					//fmt.Println("-------11-------")
				}
				for left < right && nums[right] == nums[right+1] {
					right--
					//fmt.Println("left", left)
					//fmt.Println("right", right)
					//fmt.Println("-------22-------")
				}
			} else if nums[left]+nums[right] > target {
				right--
				//fmt.Println("left", left)
				//fmt.Println("right", right)
				//fmt.Println("--------33------")
			} else if nums[left]+nums[right] < target {
				left++
				//fmt.Println("left", left)
				//fmt.Println("right", right)
				//fmt.Println("-------44-------")
			}

			//fmt.Println("------#########--------")

		}
	}

	return res

}

func sumOfTwo(nums []int, target int, left int, right int) [][]int {
	var res [][]int
	for left < right {

		var sub []int
		if nums[left]+nums[right] == target {
			fmt.Println(nums[left])
			sub = append(sub, nums[left])
			sub = append(sub, nums[right])

			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right++
			}
		} else if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		}
		if len(sub) > 0 {
			res = append(res, sub)

		}
	}

	return res
}
