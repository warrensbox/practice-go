package main

import "fmt"

/*

goal: find continuous sum in range
sol :
option 1 )use sliding window
option 2) use hashmap

lets just consider one target = 2

hash[0]=1
hash[sum+] = 1

first attempt := O(m*n)
*/

func main() {

	arr := []int{2, 1, 1, 2, 2, 1}
	left := 2
	right := 3
	continuousSum(left, right, arr)
}

func continuousSum(left, right int, arr []int) {

	hash := make(map[int]int)
	hash[0] = 1
	count, sum := 0, 0

	for target := left; target <= right; target++ {
		fmt.Println("target", target)
		for _, val := range arr {

			sum += val

			count += hash[sum-target]

			hash[sum]++
		}
	}
	fmt.Println(hash)
	fmt.Println(count)
}
