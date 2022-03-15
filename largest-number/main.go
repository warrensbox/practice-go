package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	str := make([]string, len(nums))
	for i, num := range nums {
		str[i] = strconv.Itoa(num)
	}

	sort.Slice(str, func(a, b int) bool {
		return str[a]+str[b] > str[b]+str[a]
	})

	if str[0] == "0" {
		return "0"
	}
	return strings.Join(str, "")
}
