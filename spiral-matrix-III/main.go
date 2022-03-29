package main

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {

	//1,1,2,2,3,3,4,4,
	total := rows * cols
	step := 1
	increment := 1
	res := [][]int{{rStart, cStart}}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	idx_dir := 0

	r0 := rStart
	c0 := cStart
	for len(res) < total {

		for i := 0; i < increment; i++ {
			r0 = r0 + dirs[idx_dir][0]
			c0 = c0 + dirs[idx_dir][1]
			if r0 >= 0 && c0 >= 0 && r0 < rows && c0 < cols {
				res = append(res, []int{r0, c0})
			}
		}

		idx_dir = (idx_dir + 1) % 4
		if step%2 == 0 {
			increment++
		}
		step++

	}

	return res
}
