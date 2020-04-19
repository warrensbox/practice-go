package main

import (
	"fmt"
	"math"
)

func main() {

	//rec := [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}
	//rec := [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}}
	rec := [][]int{{0, 1}, {1, 3}, {3, 3}, {4, 4}, {1, 4}, {2, 3}, {1, 0}, {3, 4}}
	output := minAreaRect(rec)

	fmt.Println(output)
}

type Points struct {
	x int
	y int
}

func minAreaRect(points [][]int) int {

	result := math.MaxInt32 //biggest value

	hashMap := make(map[Points]bool)

	var coor Points
	for _, pt := range points {
		coor.x = pt[0]
		coor.y = pt[1]
		hashMap[coor] = true
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			fmt.Println("-----")
			node1 := points[i]
			node2 := points[j]

			if node1[0] == node2[0] || node1[1] == node2[1] { //skip if not diagonal
				// fmt.Println("node1", node1)
				// fmt.Println("node2", node2)
				continue
			}

			// fmt.Println(node1)
			// fmt.Println(node2)

			node3 := Points{node1[0], node2[1]}
			node4 := Points{node2[0], node1[1]}

			if hashMap[node3] && hashMap[node4] {
				//node3-node1*node4-node1
				xAxis := node1[0] - node2[0] //node4(x)-node1(x)
				yAxis := node1[1] - node2[1]
				xDist := int(math.Abs(float64(xAxis)))
				yDist := int(math.Abs(float64(yAxis)))
				area := xDist * yDist
				if area < result {
					result = area
				}
			}
		}
	}

	fmt.Println(hashMap)

	return result
}
