package main

func main() {

}

func canJump(nums []int) bool {

	max := 0
	step := 0
	for i := 0; i < len(nums) && step >= max; i++ {

		step := i + nums[i]
		if step > max {
			max = step
		}
	}

	return step >= len(nums)-1
}
