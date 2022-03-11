package main

import "fmt"

func main() {

	fmt.Println(onBuild("A"))
	ready = make(map[string]bool)
	fmt.Println(nextBuild("D"))
	fmt.Println(nextBuild("F"))
	fmt.Println(nextBuild("E"))
}

var adjList = map[string][]string{
	"A": {"B", "C"},
	"B": {"D", "E"},
	"C": {"D", "E"},
	"G": {"F"},
	"D": nil,
	"E": nil,
	"F": nil,
}

var parentList = map[string][]string{}
var ready map[string]bool

func onBuild(node string) []string {
	//traverse graph
	res := []string{}
	queue := []string{node}
	seen := make(map[string]bool)
	parentList = make(map[string][]string)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if seen[node] {
			continue
		}
		if adjList[node] == nil {
			res = append(res, node)
		}
		seen[node] = true
		for _, dep := range adjList[node] {
			parentList[dep] = append(parentList[dep], node) //set parents
			if !seen[dep] {
				queue = append(queue, dep)
			}
		}
	}
	return res
}

func nextBuild(node string) []string {

	res := []string{}
	parents := parentList[node] //get the parent of the node
	ready[node] = true          //mark as build
	for _, parent := range parents {
		if siblingssBuildsExist(parent) { //check if sibling have been build
			res = append(res, parent)
		}
	}
	return res
}

func siblingssBuildsExist(parent string) bool {

	childrens := adjList[parent]
	count := 0
	for _, childs := range childrens {
		if ready[childs] {
			count++
		}
	}
	return len(childrens) == count
}
