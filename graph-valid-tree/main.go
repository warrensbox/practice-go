package main

func validTree(n int, edges [][]int) bool {

	//input
	//create adjacent list
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, val := range edges {
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])
	}

	// Use a set to keep track of already seen nodes to
	// avoid infinite looping.
	parent := make(map[int]int)
	parent[0] = -1
	//process

	//DFS approach
	stack := []int{0}
	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, neighbor := range graph[node] {
			if prnt, ok := parent[node]; ok && prnt == neighbor {
				continue
			}

			if _, ok := parent[neighbor]; ok { // another node has the same parent
				return false
			}
			stack = append(stack, neighbor)
			parent[neighbor] = node //who is my parent
		}
	}

	//BFS approach !!!!! COMMENT THIS BFS PART OUT
	queue := []int{0}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if prnt, ok := parent[node]; ok && prnt == neighbor {
				continue
			}

			if _, ok := parent[neighbor]; ok { // another node has the same parent
				return false
			}
			queue = append(queue, neighbor)
			parent[neighbor] = node //who is my parent
		}
	}

	// Return  output
	return len(parent) == n //all must be connect - all must be seen
}
