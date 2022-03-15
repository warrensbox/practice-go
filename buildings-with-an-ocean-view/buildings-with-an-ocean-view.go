package main

func findBuildings(heights []int) []int {

	var res []int
	max := heights[len(heights)-1]
	res = append(res, len(heights)-1)
	for i := len(heights) - 2; i >= 0; i-- {
		if heights[i] > max {
			max = heights[i]
			res = append(res, i)
		}
	}

	var res2 []int
	for len(res) > 0 {
		top := res[len(res)-1]
		res = res[:len(res)-1]
		res2 = append(res2, top)
	}

	return res2
}
