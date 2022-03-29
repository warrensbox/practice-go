package main

func combine(n int, k int) [][]int {

	arr := make([]int, n)
	res := [][]int{}
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}

	var generateComb func(index int, curr []int)
	generateComb = func(index int, curr []int) {
		if len(curr) == k {
			cp := make([]int, k)
			copy(cp, curr)
			res = append(res, cp)
		}

		for i := index; i < len(arr); i++ {
			curr = append(curr, arr[i])
			generateComb(i+1, curr)
			curr = curr[:len(curr)-1]
		}

	}
	generateComb(0, []int{})
	return res
}
