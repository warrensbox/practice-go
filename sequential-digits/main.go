package main

import "strconv"

func sequentialDigits(low int, high int) []int {
	nums := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	res := []int{}
	lowLen := len(strconv.Itoa(low))   //3
	highLen := len(strconv.Itoa(high)) //4

	//3 - 4

	for length := lowLen; length <= highLen; length++ {
		for start := 0; start < 10-length; start++ {
			byteTonum, _ := strconv.Atoi(string(nums[start : start+length]))
			if byteTonum >= low && byteTonum <= high {
				res = append(res, byteTonum)
			}
		}
	}

	return res
}
