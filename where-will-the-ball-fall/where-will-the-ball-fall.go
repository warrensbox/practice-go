package main

import "fmt"

func findBall(grid [][]int) []int {

	m := len(grid)
	n := len(grid[0])
	result := make([]int, n)

	for i := 0; i < n; i++ {
		pos := i
		fmt.Println("pos", pos)
		for j := 0; j < m; j++ {
			fmt.Println("j", j)
			dir := grid[j][pos]
			fmt.Println("dir", dir)
			nextPos := pos + dir
			fmt.Println("nextPos", nextPos)
			// fmt.Println("grid[j][nextPos]",grid[j][nextPos])
			if nextPos < 0 || nextPos >= n || grid[j][nextPos] == -dir {
				pos = -1
				break
			}
			pos += dir
			fmt.Println("pos", pos)
			fmt.Println("----")
		}
		result[i] = pos
	}

	return result
}

//where-will-the-ball-fall
