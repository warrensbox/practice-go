package main

import "fmt"

func main() {

}

func countPaths(n int, roads [][]int) int {

	graph := make(map[Node][]Node)
	//graph := make(map[int][]Node)

	for _, point := range roads {
		var node Node
		node.Val = point[1]
		node.Dist = point[2]
		graph[Node{Val: point[0]}] = append(graph[Node{Val: point[0]}], node)
		// graph[point[0]] = append(graph[point[0]],node)
	}
	fmt.Println(graph)

	var queue []Node

	queue = append(queue, Node{Val: 0})

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:]

		for _, val := range graph[node.Val] {
			queue = append(queue, val)
		}
	}

	return -1

}

type Node struct {
	Val  int
	Dist int
}
