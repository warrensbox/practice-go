package main

func allPathsSourceTarget(graph [][]int) [][]int {

	res := [][]int{}
	var dfs func(node int, arr []int)
	length := len(graph)

	dfs = func(node int, arr []int) {

		if node == length-1 {
			arr = append(arr, node)
			res = append(res, append([]int{}, arr...))
			return
		}

		for _, depends := range graph[node] {
			arr = append(arr, node)
			dfs(depends, arr)
			arr = arr[:len(arr)-1]
		}
	}

	dfs(0, []int{})

	return res

}
