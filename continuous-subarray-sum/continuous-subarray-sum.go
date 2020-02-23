package main

import (
	"fmt"
)

func main() {

	input := []int{23, 2, 4, 6, 7}
	k := 7
	output := checkSubarraySum(input, k)

	fmt.Println(output)
}

// public class Solution {
//     public boolean checkSubarraySum(int[] nums, int k) {
//         int[] sum = new int[nums.length];
//         sum[0] = nums[0];
//         for (int i = 1; i < nums.length; i++)
//             sum[i] = sum[i - 1] + nums[i];
//         for (int start = 0; start < nums.length - 1; start++) {
//             for (int end = start + 1; end < nums.length; end++) {
//                 int summ = sum[end] - sum[start] + nums[start];
//                 if (summ == k || (k != 0 && summ % k == 0))
//                     return true;
//             }
//         }
//         return false;
//     }
// }

func checkSubarraySum(nums []int, k int) bool {

	sum := make([]int, len(nums))

	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
		//fmt.Println(sum[i])
	}

	for start := 0; start < len(nums)-1; start++ {
		fmt.Println("start", start)
		for end := start + 1; end < len(nums); end++ {
			fmt.Println("end", end)
		}
		fmt.Println("-----")
	}

	return false
}

//for (int start = 0; start < nums.length - 1; start++) {
//             for (int end = start + 1; end < nums.length; end++) {
//                 int summ = sum[end] - sum[start] + nums[start];
//                 if (summ == k || (k != 0 && summ % k == 0))
//                     return true;
//             }
//         }
