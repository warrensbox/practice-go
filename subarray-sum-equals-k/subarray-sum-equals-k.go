package main

import "fmt"

func main() {

	input := []int{1, 2, 1, 2, 1}
	k := 3

	output := subarraySum(input, k)

	fmt.Println(output)

}

//sol1
func subarraySum1(nums []int, k int) int {

	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			fmt.Println("start", start)
			fmt.Println("end", end)
			sum += nums[end]
			fmt.Println("sum", sum)
			fmt.Println("-----")
			if sum == k {
				count++
			}
		}
	}

	return count
}

func subarraySum(nums []int, k int) int {

	count := 0
	sum := 0

	cache := make(map[int]int)

	cache[0] = 1

	for _, val := range nums {
		sum += val
		count += cache[sum-k]
		cache[sum]++
	}
	return count
}
