package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nums := []int{1, 2, 3, 3, 3}
	obj := Constructor(nums)
	target := 3
	param_1 := obj.Pick(target)

	fmt.Println(param_1)

}

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	var s Solution
	s.arr = nums
	return s
}

func (this *Solution) Pick(target int) int {

	total := 0
	res := -1

	for i := 0; i < len(this.arr); i++ {
		if this.arr[i] == target {
			total++
			//fmt.Println("total", total)
			//fmt.Println("rand.Intn(total)", rand.Intn(total))
			if rand.Intn(total) == 0 {
				res = i
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
