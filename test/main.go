// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func main() {

	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}

/*
goal: decode string according to num[char]
*/

type Points struct {
	x int
	y int
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	hash := make(map[int]bool, 2)
	pp := make(map[[2]int]bool, 4)

	p := [][]int{p1, p2, p3, p4}

	for i, v := range p {
		//fmt.Println(i)
		fmt.Println("v", v)
		pp[[2]int{v[0], v[1]}] = true
		//fmt.Println("pp", pp)
		for j := i + 1; j < len(p); j++ {
			pp[[2]int{v[0], v[1]}] = true
			//fmt.Println("pp2", pp)
			fmt.Println("p[j][0]", p[j][0])
			fmt.Println("v[0]", v[0])
			dist := (p[j][0]-v[0])*(p[j][0]-v[0]) + (p[j][1]-v[1])*(p[j][1]-v[1])
			fmt.Println("dist", dist)
			hash[dist] = true

			fmt.Println(hash)
			if len(hash) > 2 {
				return false
			}
		}
	}

	return true
}
