package main

import "fmt"

func main() {

	input := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}

	output := isBipartite(input)

	fmt.Println(output)

}

type Queue struct {
	data []int
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) add(e int) {
	q.data = append(q.data, e)
}

func (q *Queue) poll() int {
	e := q.data[0]
	q.data = q.data[1:]
	return e
}

func isBipartite(graph [][]int) bool {

	lengthGraph := len(graph)
	var color = make([]int, lengthGraph)

	for i := 0; i < lengthGraph; i++ {
		if color[i] == 0 {

			result := bfs(i, graph, color)
			fmt.Println(i)
			if !result {
				return false
			}
		}
	}

	return true
}

func bfs(start int, graph [][]int, color []int) bool {
	queue := Queue{}

	queue.add(start)
	color[start] = 1
	//fmt.Println(color)

	for !queue.isEmpty() {
		node := queue.poll()
		//fmt.Println("node", node)
		clr := color[node]
		//fmt.Println("clr", clr)
		//fmt.Println(color)
		for _, v := range graph[node] {
			//fmt.Println("v", v)

			if color[v] == clr {
				return false
			}

			if color[v] == 0 {
				//assign color
				if clr == 1 {
					color[v] = 2
				} else {
					color[v] = 1
				}

				queue.add(v)
			}

		}
		fmt.Println("-----")
	}

	fmt.Println("--======---")
	return true

}

//look at https://leetcode.com/problems/is-graph-bipartite/solution/
