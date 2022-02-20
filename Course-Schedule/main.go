package main

func main() {

}
func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := make(map[int][]int)

	for i, _ := range prerequisites {
		graph[i] = []int{}
	}

	for _, req := range prerequisites {
		graph[req[0]] = append(graph[req[0]], req[1])
	}

	visited := make(map[int]bool, numCourses)
	var dfs func(course int) bool
	dfs = func(course int) bool {

		if visited[course] {
			return false
		}

		if len(graph[course]) == 0 {
			return true
		}
		visited[course] = true
		for _, preq := range graph[course] {
			if !dfs(preq) {
				return false
			} //if visited, return false

		}

		delete(visited, course)
		graph[course] = []int{}
		return true
	}

	for course := 0; course < numCourses; course++ {
		if !dfs(course) {
			return false
		}
	}

	return true
}
