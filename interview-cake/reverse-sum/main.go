package main

import "fmt"

func main() {
	a := []int{2, 4, 8}
	b := []int{5, 6, 4}
	fmt.Println(reverseArrSum(a, b))
	inverseNum()
}

func reverseArrSum(a, b []int) []int {

	i, j := 0, 0
	carry := 0
	res := []int{}
	for i < len(a) || i < len(b) {

		sum := carry
		if i < len(a) {
			sum += a[i]
			i++
		}

		if j < len(b) {
			sum += b[j]
			j++
		}

		if sum > 9 {
			res = append(res, sum%10)
			carry = 1
		} else {
			res = append(res, sum)
			carry = 0
		}
	}

	if carry > 0 {
		res = append(res, carry)
	}

	return res
}

func inverseNum() {

	num := 234
	z := 0
	for num > 0 {
		y := z * 10
		x := num % 10
		z = y + x
		num = num / 10
	}
	fmt.Println(z)
}
