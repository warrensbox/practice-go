package main

func sumZero(n int) []int {

	first := 0
	last := n - 1
	res := make([]int, n)
	mag := 1
	for first <= last {

		n = n - mag
		res[first] = n
		res[last] = -1 * n

		first++
		last--
		mag = 2
	}

	return res

}
