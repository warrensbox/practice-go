package main

import (
	"fmt"
	"strings"
)

func main() {

	files := []string{
		"webpage/assets/html/a.html",
		"webpage/assets/html/b.html",
		"webpage/assets/js/c.js",
		"webpage/index.html",
	}
	constructTree2(files)
}

func constructTree2(dir []string) {

	graph := make(map[string][]string)

	for i := 0; i < len(dir); i++ {
		arr := strings.Split(dir[i], "/")
		if len(arr) > 0 {
			for j := 1; j < len(arr); j++ {
				parent := arr[j-1]
				child := arr[j]
				graph[parent] = append(graph[parent], child)
			}
		}

	}

	var dfs func(node string, indent int)
	seen := make(map[string]bool)
	dfs = func(node string, indent int) {

		if seen[node] == true {
			return
		}
		ind := strings.Repeat("-", indent)
		fmt.Println(ind + node)
		seen[node] = true
		for _, val := range graph[node] {
			dfs(val, indent+1)
		}

	}
	dfs("webpage", 0)
}
