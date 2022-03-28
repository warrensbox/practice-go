package main

import "strconv"

func minOperations(s string) int {

	count := 0
	prev, _ := strconv.Atoi(string(s[0]))
	for i := 1; i < len(s); i++ {
		currrent, _ := strconv.Atoi(string(s[i]))
		if prev^currrent != 1 {
			count++
			prev = currrent ^ 1
		} else {
			prev = currrent
		}
	}

	return Min(count, len(s)-count) //see comments
}

/*this is because
10010100
10101010 (requires 5 changes)

10010100
01010101 (requires 3 changes)

two possible ways
it's one or the other
5 or 3 (total 8)
so 8-3=5
or 8-5=3

choose the smallest
*/

func Min(x, y int) int {

	if x < y {
		return x
	}

	return y
}
