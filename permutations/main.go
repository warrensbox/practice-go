package main

func permute(nums []int) [][]int {

	res := [][]int{}
	var permSub func(first int, nums []int)

	permSub = func(first int, nums []int) {

		if first == len(nums) {
			cp := make([]int, len(nums))
			copy(cp, nums)
			res = append(res, cp)
		}

		for i := first; i < len(nums); i++ {
			nums[first], nums[i] = nums[i], nums[first]
			permSub(first+1, nums)
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	permSub(0, nums)
	return res
}
