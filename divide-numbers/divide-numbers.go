/*
goal: divide nums without multiplication, division and mod
solution: use bitwise operators
what is 10/3
- keeping track of 10-3 = 7
1
- keeping track of 7-3 = 4
2
- keeping track of 4-3 = 1
4
can I go further? No so answer = 3
*/
package main

import "math"

func divide(dividend int, divisor int) int {

	//basecase
	if dividend == 0 {
		return 0
	} else if divisor == 0 {
		return -1
	}

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	a, boolA := Abs(dividend)
	b, boolB := Abs(divisor)
	ans := division(a, b)

	if boolA && boolB {
		return ans
	} else if boolA || boolB {
		return -ans
	}
	return ans

}

func division(a, b int) int {

	res := 0
	for a-b >= 0 {
		x := 0 // 2^0
		for a-(b<<1<<x) >= 0 {
			x++
		}
		res += 1 << x
		a -= b << x
	}

	return res
}

func Abs(a int) (int, bool) {

	if a < 0 {
		return -a, true
	}

	return a, false
}
